package processor

import (
	"context"
	"fmt"
	"time"

	"github.com/cartridge-gg/starknet-indexer/ent"
	"github.com/cartridge-gg/starknet-indexer/ent/block"
	"github.com/cartridge-gg/starknet-indexer/ent/transactionreceipt"
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
	"github.com/rs/zerolog/log"
)

// Handle block persistence
type StoreBlock struct {
	BlockProcessor
}

func (p *StoreBlock) Process(ctx context.Context, rpc *jsonrpc.Client, b *types.Block) (func(tx *ent.Tx) *ProcessorError, error) {
	return func(tx *ent.Tx) *ProcessorError {
		log.Debug().Msgf("Writing block: %d", b.BlockNumber)

		if err := tx.Block.Create().
			SetID(b.BlockHash).
			SetBlockHash(b.BlockHash).
			SetBlockNumber(uint64(b.BlockNumber)).
			SetParentBlockHash(b.ParentBlockHash).
			SetStateRoot(b.NewRoot).
			SetTimestamp(time.Unix(int64(b.AcceptedTime), 0).UTC()).
			SetStatus(block.Status(b.Status)).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("block:%d", b.BlockNumber),
				Error: err,
			}
		}

		return nil
	}, nil
}

// Handle transaction persistence
type StoreTransaction struct {
	TransactionProcessor
}

func (p *StoreTransaction) Process(ctx context.Context, rpc *jsonrpc.Client, b *types.Block, txn *types.Transaction) (func(tx *ent.Tx) *ProcessorError, error) {
	return func(tx *ent.Tx) *ProcessorError {
		log.Trace().Msgf("Writing transaction: %s", txn.TransactionHash)

		if err := tx.Transaction.Create().
			SetID(txn.TransactionHash).
			SetTransactionHash(txn.TransactionHash).
			SetBlockID(b.BlockHash).
			SetContractAddress(txn.ContractAddress).
			SetEntryPointSelector(txn.EntryPointSelector).
			SetNonce(txn.Nonce).
			SetCalldata(txn.Calldata).
			SetSignature(txn.Signature).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("transaction:%s", txn.TransactionHash),
				Error: err,
			}
		}

		if err := tx.TransactionReceipt.Create().
			SetID(txn.TransactionHash).
			SetBlockID(b.BlockHash).
			SetTransactionID(txn.TransactionHash).
			SetTransactionHash(txn.TransactionHash).
			SetStatus(transactionreceipt.Status(txn.TransactionReceipt.Status)).
			SetStatusData(txn.TransactionReceipt.StatusData).
			SetMessagesSent(txn.TransactionReceipt.MessagesSent).
			SetL1OriginMessage(txn.TransactionReceipt.L1OriginMessage).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("transaction:%s", txn.TransactionHash),
				Error: err,
			}
		}

		return nil
	}, nil
}

// Handle event persistence
type StoreEvent struct {
	EventProcessor
}

func (p *StoreEvent) Process(ctx context.Context, rpc *jsonrpc.Client, b *types.Block, txn *types.Transaction, evt *Event) (func(tx *ent.Tx) *ProcessorError, error) {
	return func(tx *ent.Tx) *ProcessorError {
		log.Trace().Msgf("Writing event: %s", txn.TransactionHash)

		if err := tx.Event.Create().
			SetID(fmt.Sprintf("%s-%d", txn.TransactionHash, evt.Index)).
			SetTransactionID(txn.TransactionHash).
			SetFrom(evt.FromAddress).
			SetKeys(evt.Keys).
			SetData(evt.Data).
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("event:%s", txn.TransactionHash),
				Error: err,
			}
		}

		return nil
	}, nil
}
