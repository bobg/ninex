package main

import (
	"bytes"
	"context"
	"log"
	"ninex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func timeout(args []string) {
	fs, ethURL := newFlagSet()

	var (
		keyfile, passphrase = addKeyfilePassphrase(fs)
		contractAddr        = addContract(fs)
	)

	err := fs.Parse(args)
	must(err)

	transactOpts, err := handleKeyfilePassphrase(*keyfile, *passphrase)
	must(err)

	client, err := ethclient.Dial(*ethURL)
	must(err)

	nx, err := handleContract(*contractAddr, client)
	must(err)

	tx, err := nx.Timeout(transactOpts)
	must(err)

	ctx := context.Background()

	receipt, err := bind.WaitMined(ctx, client, tx)
	must(err)
	log.Print(receipt.String())

	a, err := abi.JSON(bytes.NewReader([]byte(ninex.NinexABI)))
	must(err)

	for _, item := range receipt.Logs {
		var ev ninex.EvWinByDefault
		if err = a.Unpack(&ev, "evWinByDefault", item.Data); err == nil {
			log.Print("Win by default:")
			log.Printf("  guesser: %x", ev.Guesser[:])
			log.Printf("  value: %s", ev.Value)
		}
	}
}
