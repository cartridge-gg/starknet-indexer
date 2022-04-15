package main

import (
	"github.com/alecthomas/kong"
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

	indexer.NewIndexer(cli.Addr)
}
