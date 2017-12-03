package main

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

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
}
