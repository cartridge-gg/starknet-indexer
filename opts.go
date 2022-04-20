package indexer

import "net/http"

type indexerOptions struct {
	client *http.Client
}

// funcIndexerOption wraps a function that modifies indexerOptions into an
// implementation of the IndexerOption interface.
type funcIndexerOption struct {
	f func(*indexerOptions)
}

func (fso *funcIndexerOption) apply(do *indexerOptions) {
	fso.f(do)
}

func newFuncIndexerOption(f func(*indexerOptions)) *funcIndexerOption {
	return &funcIndexerOption{
		f: f,
	}
}

// IndexerOption configures how we set up the connection.
type IndexerOption interface {
	apply(*indexerOptions)
}

func WithHttpClient(client http.Client) IndexerOption {
	return newFuncIndexerOption(func(o *indexerOptions) {
		o.client = &client
	})
}
