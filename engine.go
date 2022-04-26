package indexer

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
	"github.com/rs/zerolog/log"
	concurrently "github.com/tejzpr/ordered-concurrently/v3"

	"github.com/tarrencev/starknet-indexer/ent"
	"github.com/tarrencev/starknet-indexer/ent/block"
	"github.com/tarrencev/starknet-indexer/ent/transactionreceipt"
)

const parallelism = 5

type Contract struct {
	Address    string
	StartBlock uint64
	Handler    func(types.Transaction) error
}

type Config struct {
	Interval  time.Duration
	Contracts []Contract
}

type Engine struct {
	sync.Mutex
	latest   uint64
	ent      *ent.Client
	provider types.Provider
	ticker   *time.Ticker
}

func NewEngine(ctx context.Context, client *ent.Client, config Config) (*Engine, error) {
	provider, err := jsonrpc.DialContext(ctx, "http://localhost:9545")
	if err != nil {
		return nil, err
	}

	e := &Engine{
		ent:      client,
		provider: provider,
		ticker:   time.NewTicker(config.Interval),
	}

	head, err := client.Block.Query().Order(ent.Desc(block.FieldBlockNumber)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if head != nil {
		e.latest = head.BlockNumber
	}

	return e, nil
}

func (e *Engine) Start(ctx context.Context) {
	defer e.ticker.Stop()

	for {
		select {
		case <-e.ticker.C:
			e.Lock()
			e.process(ctx)
			e.Unlock()

		case <-ctx.Done():
			return
		}
	}
}

func (e *Engine) process(ctx context.Context) error {
	worker := make(chan concurrently.WorkFunction, parallelism)

	outputs := concurrently.Process(ctx, worker, &concurrently.Options{PoolSize: parallelism, OutChannelBuffer: parallelism})

	block, err := e.provider.BlockByNumber(ctx, nil, "FULL_TXN_AND_RECEIPTS")
	if err != nil {
		log.Err(err).Msg("Getting latest block number.")
		return err
	}

	head := uint64(block.BlockNumber)

	go func() {
		for i := e.latest; i < head; i++ {
			worker <- fetcher{e.provider, i}
			e.latest += 1
		}
	}()

	for output := range outputs {
		v, ok := output.Value.(response)
		if !ok {
			continue
		}

		if v.err != nil {
			log.Err(v.err).Msg("Fetching block.")
			return v.err
		}

		if err := e.write(ctx, v.block); err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) write(ctx context.Context, b *types.Block) error {
	if err := ent.WithTx(ctx, e.ent, func(tx *ent.Tx) error {
		if err := tx.Block.Create().
			SetID(b.BlockHash).
			SetBlockHash(b.BlockHash).
			SetBlockNumber(uint64(b.BlockNumber)).
			SetParentBlockHash(b.ParentBlockHash).
			SetStateRoot(b.NewRoot).
			SetTimestamp(time.Unix(int64(b.AcceptedTime), 0).UTC()).
			SetStatus(block.Status(b.Status)).
			Exec(ctx); err != nil {
			return err
		}

		for _, t := range b.Transactions {
			if err := tx.Transaction.Create().
				SetID(t.TransactionHash).
				SetTransactionHash(t.TransactionHash).
				SetBlockID(b.BlockHash).
				SetContractAddress(t.ContractAddress).
				SetEntryPointSelector(t.EntryPointSelector).
				SetNonce(t.Nonce).
				SetCalldata(t.Calldata).
				SetSignature(t.Signature).
				Exec(ctx); err != nil {
				return err
			}

			if err := tx.TransactionReceipt.Create().
				SetID(t.TransactionReceipt.TransactionHash).
				SetTransactionHash(t.TransactionReceipt.TransactionHash).
				SetStatus(transactionreceipt.Status(t.TransactionReceipt.Status)).
				SetStatusData(t.TransactionReceipt.StatusData).
				SetMessagesSent(t.TransactionReceipt.MessagesSent).
				SetL1OriginMessage(t.TransactionReceipt.L1OriginMessage).
				SetEvents(t.TransactionReceipt.Events).
				Exec(ctx); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// Create a type based on your input to the work function
type fetcher struct {
	provider    types.Provider
	blockNumber uint64
}

type response struct {
	block *types.Block
	err   error
}

// The work that needs to be performed
// The input type should implement the WorkFunction interface
func (f fetcher) Run(ctx context.Context) interface{} {
	block, err := f.provider.BlockByNumber(ctx, big.NewInt(int64(f.blockNumber)), "FULL_TXN_AND_RECEIPTS")
	return response{block, err}
}
