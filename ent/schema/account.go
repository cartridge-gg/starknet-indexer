package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

// Account defines the Account type schema.
type Account struct {
	ent.Schema
}

// Fields returns Account fields.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().Immutable().DefaultFunc(cuid.New),
		field.String("address"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
	}
}

// Edges returns Account edges.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{}
}
