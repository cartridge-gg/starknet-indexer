package main

import (
	"time"

	"github.com/alecthomas/kong"
	"github.com/dontpanicdao/caigo"
	_ "github.com/mattn/go-sqlite3"
	indexer "github.com/tarrencev/starknet-indexer"
	_ "github.com/tarrencev/starknet-indexer/ent/runtime"
)

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	indexer.New(cli.Addr, indexer.Config{
		Interval: 2 * time.Second,
		Contracts: []indexer.Contract{{
			Address:    "0x",
			StartBlock: 1000,
			Handler: func(caigo.Transaction, caigo.TransactionReceipt) error {
				// handle transaction
				return nil
			},
		}},
	})
}
