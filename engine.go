package indexer

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"time"

	"github.com/cartridge-gg/starknet-indexer/ent"
	"github.com/cartridge-gg/starknet-indexer/processor"
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
	"github.com/rs/zerolog/log"
	concurrently "github.com/tejzpr/ordered-concurrently/v3"
)

const concurrency = 5

type Contract struct {
	Address    string
	StartBlock uint64
	Handler    func(types.Transaction) error
}

type Config struct {
	Head      uint64
	Interval  time.Duration
	Contracts []Contract
}

type Engine struct {
	sync.Mutex
	client                *ent.Client
	latest                uint64
	provider              *jsonrpc.Client
	ticker                *time.Ticker
	blockProcessors       []processor.BlockProcessor
	transactionProcessors []processor.TransactionProcessor
	eventProcessors       []processor.EventProcessor
}

func NewEngine(ctx context.Context, client *ent.Client, provider *jsonrpc.Client, config Config) (*Engine, error) {
	e := &Engine{
		client:   client,
		provider: provider,
		ticker:   time.NewTicker(config.Interval),
		latest:   config.Head,
	}

	return e, nil
}

func (e *Engine) Start(ctx context.Context) {
	defer e.ticker.Stop()
	log.Info().Msg("Starting indexer.")

	for {
		select {
		case <-e.ticker.C:
			e.Lock()
			if err := e.process(ctx); err != nil {
				log.Err(err).Msg("Processing block.")
				return
			}
			e.Unlock()

		case <-ctx.Done():
			return
		}
	}
}

func (e *Engine) Register(ctx context.Context, h interface{}) error {
	switch h := h.(type) {
	case processor.BlockProcessor:
		e.Lock()
		e.blockProcessors = append(e.blockProcessors, h)
		e.Unlock()
		return nil
	case processor.TransactionProcessor:
		e.Lock()
		e.transactionProcessors = append(e.transactionProcessors, h)
		e.Unlock()
		return nil
	case processor.EventProcessor:
		e.Lock()
		e.eventProcessors = append(e.eventProcessors, h)
		e.Unlock()
		return nil
	}

	return errors.New("unsupported processor")
}

func (e *Engine) Subscribe(ctx context.Context) {

}

func (e *Engine) process(ctx context.Context) error {
	worker := make(chan concurrently.WorkFunction, concurrency)

	outputs := concurrently.Process(ctx, worker, &concurrently.Options{PoolSize: concurrency, OutChannelBuffer: concurrency})

	block, err := e.provider.BlockByNumber(ctx, nil, "FULL_TXN_AND_RECEIPTS")
	if err != nil {
		log.Err(err).Msg("Getting latest block number.")
		return err
	}

	head := uint64(block.BlockNumber)

	go func() {
		for i := e.latest; i < head; i++ {
			worker <- fetcher{e.provider, i, e.blockProcessors, e.transactionProcessors, e.eventProcessors}
		}
		close(worker)
	}()

	for output := range outputs {
		v, ok := output.Value.(response)
		if !ok {
			continue
		}

		if v.err != nil {
			log.Err(v.err).Msg("Fetching block.")
			return v.err
		}

		if err := ent.WithTx(ctx, e.client, func(tx *ent.Tx) error {
			for _, cb := range v.callbacks {
				if cb == nil {
					continue
				}

				if err := cb(tx); err != nil {
					log.Err(err).Uint64("block", uint64(v.block.BlockNumber)).Msg("Writing block.")
					return err
				}
			}

			return nil
		}); err != nil {
			return err
		}

		e.latest = uint64(v.block.BlockNumber)
	}

	return nil
}

// Create a type based on your input to the work function
type fetcher struct {
	provider              *jsonrpc.Client
	blockNumber           uint64
	blockProcessors       []processor.BlockProcessor
	transactionProcessors []processor.TransactionProcessor
	eventProcessors       []processor.EventProcessor
}

type response struct {
	block     *types.Block
	callbacks []func(*ent.Tx) error
	err       error
}

// The work that needs to be performed
// The input type should implement the WorkFunction interface
func (f fetcher) Run(ctx context.Context) interface{} {
	block, err := f.provider.BlockByNumber(ctx, big.NewInt(int64(f.blockNumber)), "FULL_TXN_AND_RECEIPTS")
	if err != nil {
		return response{block, nil, err}
	}

	var cbs []func(*ent.Tx) error
	for _, p := range f.blockProcessors {
		cb, err := p.Process(ctx, f.provider, block)
		if err != nil {
			return response{block, nil, err}
		}

		cbs = append(cbs, cb)
	}

	for _, t := range block.Transactions {
		for _, p := range f.transactionProcessors {
			cb, err := p.Process(ctx, f.provider, block, t)
			if err != nil {
				return response{block, nil, err}
			}

			cbs = append(cbs, cb)
		}

		for _, evt := range t.Events {
			for _, p := range f.eventProcessors {
				cb, err := p.Process(ctx, f.provider, block, t, evt)
				if err != nil {
					return response{block, nil, err}
				}

				cbs = append(cbs, cb)
			}
		}
	}

	return response{block, cbs, nil}
}
