package indexer

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/dontpanicdao/caigo/types"
	"github.com/rs/zerolog/log"
	concurrently "github.com/tejzpr/ordered-concurrently/v3"
)

const parallelism = 5

type WriteHandler func(ctx context.Context, block *types.Block) error

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
	latest   uint64
	provider types.Provider
	ticker   *time.Ticker
}

func NewEngine(ctx context.Context, provider types.Provider, config Config) (*Engine, error) {
	e := &Engine{
		provider: provider,
		ticker:   time.NewTicker(config.Interval),
		latest:   config.Head,
	}

	return e, nil
}

func (e *Engine) Start(ctx context.Context, writeHandler WriteHandler) {
	defer e.ticker.Stop()

	for {
		select {
		case <-e.ticker.C:
			e.Lock()
			if err := e.process(ctx, writeHandler); err != nil {
				log.Err(err).Msg("Processing block.")
			}
			e.Unlock()

		case <-ctx.Done():
			return
		}
	}
}

func (e *Engine) Subscribe(ctx context.Context) {

}

func (e *Engine) process(ctx context.Context, writeHandler WriteHandler) error {
	worker := make(chan concurrently.WorkFunction, parallelism)

	outputs := concurrently.Process(ctx, worker, &concurrently.Options{PoolSize: parallelism, OutChannelBuffer: parallelism})

	block, err := e.provider.BlockByNumber(ctx, nil, "FULL_TXN_AND_RECEIPTS")
	if err != nil {
		log.Err(err).Msg("Getting latest block number.")
		return err
	}

	head := uint64(block.BlockNumber)

	go func() {
		for i := e.latest; i < head; i++ {
			worker <- fetcher{e.provider, i}
		}
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

		if err := writeHandler(ctx, v.block); err != nil {
			log.Err(err).Msg("Writing block.")
			return err
		}

		e.latest = uint64(v.block.BlockNumber)
	}

	return nil
}

// Create a type based on your input to the work function
type fetcher struct {
	provider    types.Provider
	blockNumber uint64
}

type response struct {
	block *types.Block
	err   error
}

// The work that needs to be performed
// The input type should implement the WorkFunction interface
func (f fetcher) Run(ctx context.Context) interface{} {
	block, err := f.provider.BlockByNumber(ctx, big.NewInt(int64(f.blockNumber)), "FULL_TXN_AND_RECEIPTS")
	return response{block, err}
}
