package indexer

import (
	"context"
	"sync"
	"time"

	"github.com/dontpanicdao/caigo"
	"github.com/rs/zerolog/log"
	"github.com/tarrencev/starknet-indexer/ent"
)

type Contract struct {
	Address    string
	StartBlock uint64
	Handlers   map[string]func(caigo.Transaction, caigo.TransactionReceipt) error
}

type Config struct {
	Interval  time.Duration
	Contracts []Contract
}

type Engine struct {
	sync.Mutex
	latest    uint64
	ent       *ent.Client
	gateway   *caigo.StarknetGateway
	ticker    *time.Ticker
	contracts map[string]*Contract
}

func NewEngine(ctx context.Context, client *ent.Client, config Config) *Engine {
	gateway := caigo.NewGateway()

	e := &Engine{
		ent:     client,
		gateway: gateway,
		ticker:  time.NewTicker(config.Interval),
	}

	for _, c := range config.Contracts {
		c := c
		s, err := client.SyncState.Get(ctx, c.Address)
		if err != nil && !ent.IsNotFound(err) {
			log.Fatal().Err(err).Msg("Fetching sync state.")
		} else if err == nil {
			c.StartBlock = s.StartBlock
		}

		e.contracts[c.Address] = &c
	}

	return e
}

func (e *Engine) Start(ctx context.Context) {
	defer e.ticker.Stop()

	for {
		select {
		case <-e.ticker.C:
			e.Lock()
			block, err := e.gateway.Block(ctx, nil)
			if err != nil {
				log.Err(err).Msg("Getting latest block number.")
				continue
			}

			latest := uint64(block.BlockNumber)

			for e.latest+1 < latest {
				block, err := e.gateway.Block(ctx, nil)
				if err != nil {
					log.Err(err).Msg("Getting latest block number.")
					continue
				}

				if err := e.parse(block); err != nil {
					log.Err(err).Uint64("block_number", uint64(block.BlockNumber)).Msg("Parsing block.")
					continue
				}

				e.latest += 1
			}

			if err := e.parse(block); err != nil {
				log.Err(err).Uint64("block_number", uint64(block.BlockNumber)).Msg("Parsing block.")
				continue
			}

			e.latest = uint64(block.BlockNumber)
			e.Unlock()

		case <-ctx.Done():
			return
		}
	}
}

func (e *Engine) parse(b *caigo.Block) error {
	for i, txn := range b.Transactions {
		if c, ok := e.contracts[txn.ContractAddress]; ok {
			if h, ok := c.Handlers[txn.EntryPointSelector]; ok {
				if err := h(txn, b.TransactionReceipts[i]); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
