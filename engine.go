package indexer

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/dontpanicdao/caigo/types"
	"github.com/rs/zerolog/log"
	concurrently "github.com/tejzpr/ordered-concurrently/v3"

	"github.com/tarrencev/starknet-indexer/ent"
	"github.com/tarrencev/starknet-indexer/ent/block"
	"github.com/tarrencev/starknet-indexer/ent/transactionreceipt"
)

const parallelism = 5

type WriteHandler func(ctx context.Context, block *types.Block) error

type Contract struct {
	Address    string
	StartBlock uint64
	Handler    func(types.Transaction) error
}

type Config struct {
	Head         uint64
	Interval     time.Duration
	Contracts    []Contract
	WriteHandler *WriteHandler
}

type Engine struct {
	sync.Mutex
	latest       uint64
	ent          *ent.Client
	provider     types.Provider
	ticker       *time.Ticker
	writeHandler *WriteHandler
}

func NewEngine(ctx context.Context, client *ent.Client, provider types.Provider, config Config) (*Engine, error) {
	e := &Engine{
		ent:          client,
		provider:     provider,
		ticker:       time.NewTicker(config.Interval),
		latest:       config.Head,
		writeHandler: config.WriteHandler,
	}

	return e, nil
}

func (e *Engine) Start(ctx context.Context) {
	defer e.ticker.Stop()

	for {
		select {
		case <-e.ticker.C:
			e.Lock()
			if err := e.process(ctx); err != nil {
				log.Err(err).Msg("Processing block.")
			}
			e.Unlock()

		case <-ctx.Done():
			return
		}
	}
}

func (e *Engine) Subscribe(ctx context.Context) {

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

		if e.writeHandler != nil {
			if err := (*e.writeHandler)(ctx, v.block); err != nil {
				log.Err(err).Msg("Writing block.")
				return err
			}
		} else {
			if err := e.write(ctx, v.block); err != nil {
				log.Err(err).Msg("Writing block.")
				return err
			}
		}
	}

	return nil
}

func (e *Engine) write(ctx context.Context, b *types.Block) error {
	log.Info().Msgf("Processing block: %d", b.BlockNumber)

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
				SetID(t.TransactionHash).
				SetBlockID(b.BlockHash).
				SetTransactionID(t.TransactionReceipt.TransactionHash).
				SetTransactionHash(t.TransactionReceipt.TransactionHash).
				SetStatus(transactionreceipt.Status(t.TransactionReceipt.Status)).
				SetStatusData(t.TransactionReceipt.StatusData).
				SetMessagesSent(t.TransactionReceipt.MessagesSent).
				SetL1OriginMessage(t.TransactionReceipt.L1OriginMessage).
				Exec(ctx); err != nil {
				return err
			}

			for i, e := range t.TransactionReceipt.Events {
				for j, k := range e.Keys {
					if err := tx.Event.Create().
						SetID(fmt.Sprintf("%s-%d-%d", t.TransactionHash, i, j)).
						SetTransactionID(t.TransactionHash).
						SetFrom(e.FromAddress).
						SetKey(k).
						SetValue(e.Values[j]).
						Exec(ctx); err != nil {
						return err
					}
				}
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
