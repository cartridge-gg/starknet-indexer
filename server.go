package indexer

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tarrencev/starknet-indexer/ent"
)

func NewIndexer() {
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatal().Err(err).Msg("opening ent client")
	}
	ctx := context.Background()
	e := NewEngine(ctx, client, Config{
		Interval: 1 * time.Second,
	})
	e.Start(ctx)
}
