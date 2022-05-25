package indexer

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tarrencev/starknet-indexer/ent"
	"github.com/tarrencev/starknet-indexer/gql"
)

func (r *balanceResolver) Balance(ctx context.Context, obj *ent.Balance) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *tokenResolver) Tokenid(ctx context.Context, obj *ent.Token) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Balance returns gql.BalanceResolver implementation.
func (r *Resolver) Balance() gql.BalanceResolver { return &balanceResolver{r} }

// Token returns gql.TokenResolver implementation.
func (r *Resolver) Token() gql.TokenResolver { return &tokenResolver{r} }

type balanceResolver struct{ *Resolver }
type tokenResolver struct{ *Resolver }
