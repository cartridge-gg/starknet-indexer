package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/dontpanicdao/caigo/types"
)

var FeltSchemaType = map[string]string{
	dialect.Postgres: "numeric",
}

// Event defines the Event type schema.
type Event struct {
	ent.Schema
}

// Fields returns Event fields.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable(),
		field.String("from"),
		field.JSON("keys", []*types.Felt{}).Annotations(entgql.Type("[Felt]")),
		field.JSON("data", []*types.Felt{}).Annotations(entgql.Type("[Felt]")),
	}
}

// Edges returns Event edges.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("transaction", Transaction.Type).Ref("events").
			Unique(),
	}
}

// Annotations returns Event annotations.
func (Event) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
