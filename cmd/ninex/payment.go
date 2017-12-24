package main

import (
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func payment(args []string) {
	fs, ethURL := newFlagSet()

	var (
		contractAddr = addContract(fs)
		accountHex   = fs.String("account", "", "account address (hex)")
	)

	err := fs.Parse(args)
	must(err)

	client, err := ethclient.Dial(*ethURL)
	must(err)

	nx, err := handleContract(*contractAddr, client)
	must(err)

	var account common.Address
	_, err = hex.Decode(account[:], []byte(*accountHex))
	must(err)

	callOpts := new(bind.CallOpts)

	amount, err := nx.MPayments(callOpts, account)
	must(err)

	log.Printf("%s", amount)
}
