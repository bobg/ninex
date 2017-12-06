package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"log"
	"math/big"
	"ninex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// This must track the definition of evGuess in ninex.sol
type EvGuess struct {
	Guesser    common.Address
	Commitment [32]byte
	Digits     []byte
	Stake      *big.Int
}

func guess(args []string) {
	fs, ethURL := newFlagSet()

	var (
		keyfile, passphrase = addKeyfilePassphrase(fs)
		contractAddr        = addContract(fs)
		wei                 = addValue(fs)
		commitmentHex       = fs.String("commitment", "", "commitment against which the guess is made")
		digits              = fs.String("digits", "", "digits of the guess")
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

	commitmentBytes, err := hex.DecodeString(*commitmentHex)
	must(err)

	var commitment [32]byte
	copy(commitment[:], commitmentBytes)

	transactOpts.GasLimit = big.NewInt(1000000)

	tx, err := nx.Guess(transactOpts, commitment, []byte(*digits))
	must(err)

	ctx := context.Background()

	receipt, err := bind.WaitMined(ctx, client, tx)
	must(err)
	log.Print(receipt.String())

	a, err := abi.JSON(bytes.NewReader([]byte(ninex.NinexABI)))
	must(err)

	for _, item := range receipt.Logs {
		var ev EvGuess
		if err = a.Unpack(&ev, "evGuess", item.Data); err == nil {
			log.Print("Guess:")
			log.Printf("  guesser: %x", ev.Guesser[:])
			log.Printf("  commitment: %x", ev.Commitment[:])
			log.Printf("  digits: %s", string(ev.Digits))
			log.Printf("  stake: %s", ev.Stake)
		}
	}
}
