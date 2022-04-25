package indexer

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tarrencev/starknet-indexer/ent"
	"github.com/tarrencev/starknet-indexer/gql"
)

func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

func (r *queryResolver) Blocks(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.BlockOrder, where *gql.BlockWhereInput) (*ent.BlockConnection, error) {
	return r.client.Block.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithBlockOrder(orderBy),
		)
}

func (r *queryResolver) Transactions(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *gql.TransactionWhereInput) (*ent.TransactionConnection, error) {
	return r.client.Transaction.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTransactionOrder(nil),
		)
}

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }