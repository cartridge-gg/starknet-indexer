package indexer

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
	"github.com/rs/zerolog/log"
	"github.com/tarrencev/starknet-indexer/ent"
	"github.com/tarrencev/starknet-indexer/ent/block"
	"github.com/tarrencev/starknet-indexer/ent/transactionreceipt"
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

	e, err := NewEngine(ctx, provider, Config{
		Head:     head.BlockNumber,
		Interval: 1 * time.Second,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Initializing engine.")
	}

	go e.Start(ctx, func(ctx context.Context, b *types.Block) error {
		log.Info().Msgf("Processing block: %d", b.BlockNumber)

		if err := ent.WithTx(ctx, e.ent, func(tx *ent.Tx) error {
			if err := tx.Block.Create().
				SetID(b.BlockHash).
				SetBlockHash(b.BlockHash).
				SetBlockNumber(uint64(b.BlockNumber)).
				SetParentBlockHash(b.ParentBlockHash).
				SetStateRoot(b.NewRoot).
				SetTimestamp(time.Unix(int64(b.AcceptedTime), 0).UTC()).
				SetStatus(block.Status(b.Status)).
				Exec(ctx); err != nil {
				return err
			}

			for _, t := range b.Transactions {
				if err := tx.Transaction.Create().
					SetID(t.TransactionHash).
					SetTransactionHash(t.TransactionHash).
					SetBlockID(b.BlockHash).
					SetContractAddress(t.ContractAddress).
					SetEntryPointSelector(t.EntryPointSelector).
					SetNonce(t.Nonce).
					SetCalldata(t.Calldata).
					SetSignature(t.Signature).
					Exec(ctx); err != nil {
					return err
				}

				if err := tx.TransactionReceipt.Create().
					SetID(t.TransactionHash).
					SetBlockID(b.BlockHash).
					SetTransactionID(t.TransactionReceipt.TransactionHash).
					SetTransactionHash(t.TransactionReceipt.TransactionHash).
					SetStatus(transactionreceipt.Status(t.TransactionReceipt.Status)).
					SetStatusData(t.TransactionReceipt.StatusData).
					SetMessagesSent(t.TransactionReceipt.MessagesSent).
					SetL1OriginMessage(t.TransactionReceipt.L1OriginMessage).
					Exec(ctx); err != nil {
					return err
				}

				for i, e := range t.TransactionReceipt.Events {
					for j, k := range e.Keys {
						if err := tx.Event.Create().
							SetID(fmt.Sprintf("%s-%d-%d", t.TransactionHash, i, j)).
							SetTransactionID(t.TransactionHash).
							SetFrom(e.FromAddress).
							SetKey(k).
							SetValue(e.Values[j]).
							Exec(ctx); err != nil {
							return err
						}
					}
				}
			}

			return nil
		}); err != nil {
			return err
		}

		return nil
	})

	log.Info().Str("address", addr).Msg("listening on")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Err(err).Msg("http server terminated")
	}
}
