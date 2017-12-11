package main

import (
	"bytes"
	"context"
	"log"
	"math/big"
	"ninex"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

	transactOpts.GasLimit = big.NewInt(5000000) // xxx

	tx, err := nx.Reveal(transactOpts, []byte(*preimage))
	must(err)

	ctx := context.Background()

	receipt, err := bind.WaitMined(ctx, client, tx)
	must(err)
	log.Print(receipt.String())

	a, err := abi.JSON(bytes.NewReader([]byte(ninex.NinexABI)))
	must(err)

	for _, item := range receipt.Logs {
		var (
			evWin  ninex.EvWin
			evLose ninex.EvLose
		)
		if err = a.Unpack(&evWin, "evWin", item.Data); err == nil {
			log.Print("Win:")
			log.Printf("  guesser: %x", evWin.Guesser[:])
			log.Printf("  value: %s", evWin.Value)
			log.Printf("  digits: %s", string(evWin.Digits))
			log.Printf("  commitment: %x", evWin.Commitment[:])
		}
		if err = a.Unpack(&evLose, "evLose", item.Data); err == nil {
			log.Print("Lose:")
			log.Printf("  guesser: %x", evLose.Guesser[:])
			log.Printf("  digits: %s", string(evLose.Digits))
			log.Printf("  commitment: %x", evLose.Commitment[:])
		}
	}
}
