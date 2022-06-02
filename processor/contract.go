package processor

import (
	"context"
	"fmt"

	"github.com/cartridge-gg/starknet-indexer/ent"
	"github.com/cartridge-gg/starknet-indexer/ent/contract"
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
	"github.com/rs/zerolog/log"
)

type MatchableContract interface {
	Address() string
	Type() string
	Match(ctx context.Context, provider *jsonrpc.Client) bool
}

// Unknown
type UnknownContract struct {
	MatchableContract
	address string
}

func NewUnknownContract(address string) *UnknownContract {
	return &UnknownContract{address: address}
}

func (c *UnknownContract) Address() string {
	return c.address
}

func (c *UnknownContract) Type() string {
	return "UNKNOWN"
}

func Match(ctx context.Context, provider *jsonrpc.Client, address string) MatchableContract {
	if c := NewERC20Contract(address); c.Match(ctx, provider) {
		return c
	}

	if c := NewERC721Contract(address); c.Match(ctx, provider) {
		return c
	}

	return NewUnknownContract(address)
}

type StoreContract struct {
	TransactionProcessor
}

// Handle contract persistence
func (p *StoreContract) Process(ctx context.Context, rpc *jsonrpc.Client, b *types.Block, txn *types.Transaction) (func(tx *ent.Tx) *ProcessorError, error) {
	// txn "type" field empty. check if call data & entry point selector are empty for now
	// to know if txn is of deploy type
	if txn.EntryPointSelector != "" || txn.Status == types.REJECTED.String() {
		return nil, nil
	}

	m := Match(ctx, rpc, txn.ContractAddress)
	return func(tx *ent.Tx) *ProcessorError {
		log.Debug().Msgf("Writing matched contract: %s:%s", m.Type(), m.Address())

		if err := tx.Contract.Create().
			SetID(m.Address()).
			SetType(contract.Type(m.Type())).
			OnConflictColumns("id").
			DoNothing().
			Exec(ctx); err != nil {
			return &ProcessorError{
				Scope: fmt.Sprintf("contract:%s", m.Address()),
				Error: err,
			}
		}

		return nil
	}, nil
}
