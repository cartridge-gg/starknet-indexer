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
	"github.com/dontpanicdao/caigo"
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
	"github.com/rs/zerolog/log"
	"github.com/tarrencev/starknet-indexer/ent"
	"github.com/tarrencev/starknet-indexer/ent/block"
	"github.com/tarrencev/starknet-indexer/ent/contract"
	"github.com/tarrencev/starknet-indexer/ent/transactionreceipt"
	"github.com/tarrencev/starknet-indexer/processor"
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
		n = head.BlockNumber
	}

	e, err := NewEngine(ctx, provider, Config{
		Head:     n,
		Interval: 1 * time.Second,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Initializing engine.")
	}

	go e.Start(ctx, func(ctx context.Context, b *types.Block) (func() error, error) {
		log.Info().Msgf("Processing block: %d", b.BlockNumber)
		var matches []processor.MatchableContract
		var balanceUpdates []BalanceUpdate

		for _, tx := range b.Transactions {
			// check if transaction emitted a transfer event
			for _, event := range tx.Events {
				if len(event.Keys) == 0 || event.Keys[0].Cmp(caigo.GetSelectorFromName("Transfer")) != 0 {
					continue
				}

				// check if contract indexed
				balanceUpdate := BalanceUpdate{
					Event: event,
				}
				contract, _ := client.Contract.Query().Where(contract.IDEQ(tx.ContractAddress)).Only(ctx)
				// index contract if not indexed
				if contract == nil {
					matched := processor.Match(ctx, tx.ContractAddress, provider)
					matches = append(matches, matched)

					if matched.Type() != "UNKNOWN" {
						balanceUpdate.ContractAddress = tx.ContractAddress
						balanceUpdate.ContractType = contract.Type.String()
						balanceUpdates = append(balanceUpdates, balanceUpdate)
						log.Info().Msgf("Matched contract at %s with %s", tx.ContractAddress, matched.Type())
					}
				} else if contract.Type.String() != "UNKNOWN" {
					balanceUpdate.ContractAddress = tx.ContractAddress
					balanceUpdate.ContractType = contract.Type.String()
					balanceUpdates = append(balanceUpdates, balanceUpdate)
				}
			}
		}

		return func() error {
			log.Info().Msgf("Writing block: %d", b.BlockNumber)

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
				}

				for _, m := range matches {
					if err := tx.Contract.Create().
						SetID(m.Address()).
						SetType(contract.Type(m.Type())).
						Exec(ctx); err != nil {
						return err
					}
				}

				for _, u := range balanceUpdates {
					switch u.ContractType {
					case "ERC20":
						senderBalance, _ := tx.Balance.Get(ctx, fmt.Sprintf("%s:%s", u.Event.Data[0], u.ContractAddress))
						if senderBalance == nil {
							if err := tx.Balance.Create().
								SetID(fmt.Sprintf("%s:%s", u.Event.Data[0].Hex(), u.ContractAddress)).
								SetAccountID(u.Event.Data[0].Hex()).
								SetContractID(u.ContractAddress).
								Exec(ctx); err != nil {
								return err
							}
						} else {
							if err := senderBalance.Update().
								AddBalance(-u.Event.Data[2].Int64()).
								Exec(ctx); err != nil {
								return err
							}
						}

						receiverBalance, _ := tx.Balance.Get(ctx, fmt.Sprintf("%s:%s", u.Event.Data[1], u.ContractAddress))
						if receiverBalance == nil {
							if err := tx.Balance.Create().
								SetID(fmt.Sprintf("%s:%s", u.Event.Data[0].Hex(), u.ContractAddress)).
								SetAccountID(u.Event.Data[0].Hex()).
								SetContractID(u.ContractAddress).
								SetBalance(u.Event.Data[2].Uint64()).
								Exec(ctx); err != nil {
								return err
							}
						} else {
							if err := receiverBalance.Update().
								AddBalance(u.Event.Data[2].Int64()).
								Exec(ctx); err != nil {
								return err
							}
						}
					case "ERC721":
						token, _ := tx.Token.Get(ctx, fmt.Sprintf("$%s:%s", u.ContractAddress, u.Event.Data[2].String()))
						if token == nil {
							if err := tx.Token.Create().
								SetID(fmt.Sprintf("$%s:%s", u.ContractAddress, u.Event.Data[2].String())).
								SetOwnerID(u.Event.Data[1].Hex()).
								SetContractID(u.ContractAddress).
								SetTokenId(u.Event.Data[2].Uint64()).
								Exec(ctx); err != nil {
								return err
							}
						} else {
							if err := token.Update().
								SetOwnerID(u.Event.Data[1].Hex()).
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
		}, nil
	})

	log.Info().Str("address", addr).Msg("listening on")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Err(err).Msg("http server terminated")
	}
}
