package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var bigTen = big.NewInt(10)

func setCommitment(args []string) {
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

	var preimage [78]byte
	for i := 0; i < len(preimage); i++ {
		n, err := rand.Int(rand.Reader, bigTen)
		must(err)
		preimage[i] = '0' + byte(n.Uint64())
	}
	log.Printf("preimage: %s", string(preimage[:]))

	transactOpts.GasLimit = big.NewInt(1000000)

	commitment := sha256.Sum256(preimage[:])
	tx, err := nx.SetCommitment(transactOpts, commitment)
	must(err)

	ctx := context.Background()

	receipt, err := bind.WaitMined(ctx, client, tx)
	must(err)
	log.Print(receipt.String())
}
