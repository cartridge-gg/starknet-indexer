package indexer

import (
	"context"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/zerolog/log"
	"github.com/tarrencev/starknet-indexer/ent"
)

func New(addr string, drv *sql.Driver, config Config, opts ...IndexerOption) {
	iopts := indexerOptions{
		debug:  false,
		client: http.DefaultClient,
	}

	for _, opt := range opts {
		opt.apply(&iopts)
	}

	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(
		context.Background(),
	); err != nil {
		log.Fatal().Err(err).Msg("Running schema migration")
	}

	srv := handler.NewDefaultServer(NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	if iopts.debug {
		srv.Use(&debug.Tracer{})
	}

	http.Handle("/",
		playground.Handler("Starknet Indexer", "/query"),
	)
	http.Handle("/query", srv)

	ctx := context.Background()
	e, err := NewEngine(ctx, client, Config{
		Interval: 1 * time.Second,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Initializing engine.")
	}

	go e.Start(ctx)

	log.Info().Str("address", addr).Msg("listening on")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Err(err).Msg("http server terminated")
	}
}
