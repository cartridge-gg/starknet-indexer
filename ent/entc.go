//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithWhereFilters(true),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../entgql.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
	}
	err = entc.Generate("./schema", &gen.Config{
		IDType: &field.TypeInfo{Type: field.TypeString},
		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
	}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}

	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}
	err = api.Generate(cfg, api.PrependPlugin(ex.CreatePlugin()))
	if err != nil {
		log.Fatalf("running gqlgen: %v", err)
	}
}
