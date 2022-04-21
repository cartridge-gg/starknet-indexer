package indexer

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/dontpanicdao/caigo"
	"github.com/rs/zerolog/log"
	"github.com/tarrencev/starknet-indexer/ent"
	"github.com/tarrencev/starknet-indexer/ent/block"
	"github.com/tarrencev/starknet-indexer/ent/schema"
	"github.com/tarrencev/starknet-indexer/ent/transaction"
)

type Contract struct {
	Address    string
	StartBlock uint64
	Handler    func(caigo.Transaction, caigo.TransactionReceipt) error
}

type Config struct {
	Interval  time.Duration
	Contracts []Contract
}

type Engine struct {
	sync.Mutex
	latest    uint64
	ent       *ent.Client
	gateway   *caigo.StarknetGateway
	ticker    *time.Ticker
	contracts map[string]*Contract
}

func NewEngine(ctx context.Context, client *ent.Client, config Config) *Engine {
	gateway := caigo.NewGateway()

	return &Engine{
		ent:     client,
		gateway: gateway,
		ticker:  time.NewTicker(config.Interval),
	}
}

func (e *Engine) Start(ctx context.Context) {
	defer e.ticker.Stop()

	for {
		select {
		case <-e.ticker.C:
			e.Lock()
			block, err := e.gateway.Block(ctx, nil)
			if err != nil {
				log.Err(err).Msg("Getting latest block number.")
				continue
			}

			latest := uint64(block.BlockNumber)

			for e.latest+1 < latest {
				block, err := e.gateway.Block(ctx, &caigo.BlockOptions{BlockNumber: int(e.latest)})
				if err != nil {
					log.Err(err).Msg("Getting latest block number.")
					continue
				}

				if err := e.parse(ctx, block); err != nil {
					log.Err(err).Uint64("block_number", uint64(block.BlockNumber)).Msg("Parsing block.")
					continue
				}

				e.latest += 1
			}

			if err := e.parse(ctx, block); err != nil {
				log.Err(err).Uint64("block_number", uint64(block.BlockNumber)).Msg("Parsing block.")
				continue
			}

			e.latest = uint64(block.BlockNumber)
			e.Unlock()

		case <-ctx.Done():
			return
		}
	}
}

func (e *Engine) parse(ctx context.Context, b *caigo.Block) error {
	if err := ent.WithTx(ctx, e.ent, func(tx *ent.Tx) error {
		if err := tx.Block.Create().
			SetID(b.BlockHash).
			SetBlockHash(b.BlockHash).
			SetBlockNumber(uint64(b.BlockNumber)).
			SetParentBlockHash(b.ParentBlockHash).
			SetStateRoot(b.StateRoot).
			SetTimestamp(time.Unix(int64(b.Timestamp), 0).UTC()).
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
				SetEntryPointType(t.EntryPointType).
				SetNonce(t.Nonce).
				SetType(transaction.Type(t.Type)).
				SetCalldata(t.Calldata).
				SetSignature(t.Signature).
				Exec(ctx); err != nil {
				return err
			}
		}

		for _, t := range b.TransactionReceipts {
			events, err := json.Marshal(t.Events)
			if err != nil {
				return err
			}

			l2ToL1Messages, err := json.Marshal(t.L2ToL1Messages)
			if err != nil {
				return err
			}

			executionResources := schema.ExecutionResources{
				NSteps:       uint64(t.ExecutionResources.NSteps),
				NMemoryHoles: uint64(t.ExecutionResources.NMemoryHoles),
				BuiltinInstanceCounter: schema.BuiltinInstanceCounter{
					PedersenBuiltin:   uint64(t.ExecutionResources.BuiltinInstanceCounter.PedersenBuiltin),
					RangeCheckBuiltin: uint64(t.ExecutionResources.BuiltinInstanceCounter.RangeCheckBuiltin),
					BitwiseBuiltin:    uint64(t.ExecutionResources.BuiltinInstanceCounter.BitwiseBuiltin),
					OutputBuiltin:     uint64(t.ExecutionResources.BuiltinInstanceCounter.OutputBuiltin),
					EcdsaBuiltin:      uint64(t.ExecutionResources.BuiltinInstanceCounter.EcdsaBuiltin),
					EcOpBuiltin:       uint64(t.ExecutionResources.BuiltinInstanceCounter.EcOpBuiltin),
				},
			}

			if err := tx.TransactionReceipt.Create().
				SetID(t.TransactionHash).
				SetTransactionHash(t.TransactionHash).
				SetBlockID(b.BlockHash).
				SetTransactionID(t.TransactionHash).
				SetTransactionIndex(int32(t.TransactionIndex)).
				SetL1ToL2ConsumedMessage(t.L1ToL2ConsumedMessage).
				SetExecutionResources(executionResources).
				SetEvents(events).
				SetL2ToL1Messages(l2ToL1Messages).
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
