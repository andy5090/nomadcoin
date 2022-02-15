package main

import (
	"github.com/andy5090/nomadcoin/blockchain"
	"github.com/andy5090/nomadcoin/cli"
	"github.com/andy5090/nomadcoin/db"
)

func main() {
	defer db.Close()
	blockchain.Blockchain()
	cli.Start()
}
