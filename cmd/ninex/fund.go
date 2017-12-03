package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func fund(args []string) {
	fs, ethURL := newFlagSet()

	var (
		keyfile, passphrase = addKeyfilePassphrase(fs)
		contractAddr        = addContract(fs)
		wei                 = addValue(fs)
	)

	err := fs.Parse(args)
	must(err)

	transactOpts, err := handleKeyfilePassphrase(*keyfile, *passphrase)
	must(err)

	client, err := ethclient.Dial(*ethURL)
	must(err)

	nx, err := handleContract(*contractAddr, client)
	must(err)

	transactOpts.Value = handleValue(*wei)

	tx, err := nx.Fund(transactOpts)
	must(err)

	ctx := context.Background()

	receipt, err := bind.WaitMined(ctx, client, tx)
	must(err)
	log.Print(receipt.String())
}
