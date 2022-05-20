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
	"github.com/tarrencev/starknet-indexer/ent/contract"
	"github.com/tarrencev/starknet-indexer/ent/transactionreceipt"
	"github.com/tarrencev/starknet-indexer/processor"
)

func New(addr string, drv *sql.Driver, provider jsonrpc.Client, config Config, opts ...IndexerOption) {
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
		n = head.BlockNumber
	}

	e, err := NewEngine(ctx, provider, Config{
		Head:     n,
		Interval: 1 * time.Second,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Initializing engine.")
	}

	go e.Start(ctx, func(ctx context.Context, b *types.Block) error {
		log.Info().Msgf("Processing block: %d", b.BlockNumber)
		if err := ent.WithTx(ctx, client, func(tx *ent.Tx) error {
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
					SetTransactionID(t.TransactionHash).
					SetTransactionHash(t.TransactionHash).
					SetStatus(transactionreceipt.Status(t.TransactionReceipt.Status)).
					SetStatusData(t.TransactionReceipt.StatusData).
					SetMessagesSent(t.TransactionReceipt.MessagesSent).
					SetL1OriginMessage(t.TransactionReceipt.L1OriginMessage).
					Exec(ctx); err != nil {
					return err
				}

				for i, e := range t.TransactionReceipt.Events {
					if err := tx.Event.Create().
						SetID(fmt.Sprintf("%s-%d", t.TransactionHash, i)).
						SetTransactionID(t.TransactionHash).
						SetFrom(e.FromAddress).
						SetKeys(e.Keys).
						SetData(e.Data).
						Exec(ctx); err != nil {
						return err
					}
				}

				if t.Type == "deploy" && t.Status != "REJECTED" {
					contractCode, err := provider.CodeAt(ctx, t.ContractAddress)
					if err == nil {
						matchedContract := processor.Match(ctx, t.ContractAddress, contractCode, provider)
						if matchedContract != nil {
							if err := tx.Contract.Create().
								SetID(t.ContractAddress).
								SetType(contract.Type(matchedContract.Type())).
								Exec(ctx); err != nil {
								return err
							}
						} else {
							if err := tx.Contract.Create().
								SetID(t.ContractAddress).
								SetType(contract.TypeUNKNOWN).
								Exec(ctx); err != nil {
								return err
							}
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
