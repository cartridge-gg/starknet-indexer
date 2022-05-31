package processor

import (
	"context"

	"github.com/cartridge-gg/starknet-indexer/ent"
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
)

type BlockProcessor interface {
	Process(context.Context, *jsonrpc.Client, *types.Block) (func(*ent.Tx) error, error)
}
type TransactionProcessor interface {
	Process(context.Context, *jsonrpc.Client, *types.Block, *types.Transaction) (func(*ent.Tx) error, error)
}
type EventProcessor interface {
	Process(context.Context, *jsonrpc.Client, *types.Block, *types.Transaction, *Event) (func(*ent.Tx) error, error)
}

type Event struct {
	*types.Event
	Index uint64
}
