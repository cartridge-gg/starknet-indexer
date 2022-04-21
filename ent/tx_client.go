package ent

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func WithTx(ctx context.Context, client *Client, fn func(tx *Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			if err := tx.Rollback(); err != nil {
				log.Err(err).Msg("Rolling back txn in recover.")
			}
			log.Error().Msgf("Rolled back txn: %v", v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
