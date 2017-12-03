package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func reveal(args []string) {
	fs, ethURL := newFlagSet()

	var (
		keyfile, passphrase = addKeyfilePassphrase(fs)
		contractAddr        = addContract(fs)
		preimage            = fs.String("preimage", "", "preimage")
	)

	err := fs.Parse(args)
	must(err)

	transactOpts, err := handleKeyfilePassphrase(*keyfile, *passphrase)
	must(err)

	client, err := ethclient.Dial(*ethURL)
	must(err)

	nx, err := handleContract(*contractAddr, client)
	must(err)

	transactOpts.GasLimit = big.NewInt(1000000) // xxx

	tx, err := nx.Reveal(transactOpts, []byte(*preimage))
	must(err)

	ctx := context.Background()

	receipt, err := bind.WaitMined(ctx, client, tx)
	must(err)
	log.Print(receipt.String())
}
