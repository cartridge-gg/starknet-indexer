package indexer

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/cartridge-gg/starknet-indexer/ent"
	"github.com/cartridge-gg/starknet-indexer/gql"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return gql.NewExecutableSchema(gql.Config{
		Resolvers: &Resolver{client},
	})
}
