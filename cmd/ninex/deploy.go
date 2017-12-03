package main

import (
	"context"
	"log"
	"ninex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func deploy(args []string) {
	fs, ethURL := newFlagSet()
	keyfile, passphrase := addKeyfilePassphrase(fs)

	err := fs.Parse(args)
	must(err)

	transactOpts, err := handleKeyfilePassphrase(*keyfile, *passphrase)
	must(err)

	client, err := ethclient.Dial(*ethURL)
	must(err)

	address, tx, _, err := ninex.DeployNinex(transactOpts, client)
	must(err)

	log.Printf("deployed ninex contract at %x", address[:])

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	must(err)

	log.Print(receipt.String())
}
