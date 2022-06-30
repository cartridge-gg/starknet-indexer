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
	"github.com/cartridge-gg/starknet-indexer/ent"
	"github.com/cartridge-gg/starknet-indexer/ent/block"
	"github.com/cartridge-gg/starknet-indexer/processor"
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
	"github.com/rs/zerolog/log"
)

type BalanceUpdate struct {
	// https://github.com/OpenZeppelin/cairo-contracts/blob/main/src/openzeppelin/token/erc20/library.cairo#L20
	// https://github.com/OpenZeppelin/cairo-contracts/blob/main/src/openzeppelin/token/erc721/library.cairo#L30
	Event           *types.Event
	ContractAddress string
	ContractType    string
}

func New(addr string, drv *sql.Driver, provider *jsonrpc.Client, config Config, opts ...IndexerOption) {
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

	var n uint64
	head, err := client.Block.Query().Order(ent.Desc(block.FieldBlockNumber)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Fatal().Err(err).Msg("Getting head block")
	} else if head != nil {
		n = head.BlockNumber + 1
	}

	e, err := NewEngine(ctx, client, provider, Config{
		Head:     n,
		Interval: 1 * time.Second,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Initializing engine.")
	}

	if err := e.Register(ctx, new(processor.StoreBlock)); err != nil {
		log.Fatal().Err(err).Msg("Registering processor.")
	}
	if err := e.Register(ctx, new(processor.StoreTransaction)); err != nil {
		log.Fatal().Err(err).Msg("Registering processor.")
	}
	if err := e.Register(ctx, new(processor.StoreEvent)); err != nil {
		log.Fatal().Err(err).Msg("Registering processor.")
	}
	if err := e.Register(ctx, new(processor.StoreContract)); err != nil {
		log.Fatal().Err(err).Msg("Registering processor.")
	}
	if err := e.Register(ctx, new(processor.ERC20Contract)); err != nil {
		log.Fatal().Err(err).Msg("Registering processor.")
	}
	if err := e.Register(ctx, new(processor.ERC721Contract)); err != nil {
		log.Fatal().Err(err).Msg("Registering processor.")
	}
	if err := e.Register(ctx, new(processor.ERC1155Contract)); err != nil {
		log.Fatal().Err(err).Msg("Registering processor.")
	}

	go e.Start(ctx)

	log.Info().Str("address", addr).Msg("listening on")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Err(err).Msg("http server terminated")
	}
}
