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
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/rs/zerolog/log"
	"github.com/tarrencev/starknet-indexer/ent"
	"github.com/tarrencev/starknet-indexer/ent/block"
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

	head, err := client.Block.Query().Order(ent.Desc(block.FieldBlockNumber)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Fatal().Err(err).Msg("Getting head block")
	}

	provider, err := jsonrpc.DialContext(ctx, "http://localhost:9545")
	if err != nil {
		log.Fatal().Err(err).Msg("Dialing provider")
	}

	e, err := NewEngine(ctx, client, provider, Config{
		Head:     head.BlockNumber,
		Interval: 1 * time.Second,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Initializing engine.")
	}

	go e.Start(ctx, nil)

	log.Info().Str("address", addr).Msg("listening on")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Err(err).Msg("http server terminated")
	}
}
