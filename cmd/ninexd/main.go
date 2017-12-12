package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"
	"ninex"
	"os"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	nx            *ninex.Ninex
	adminSecret   [16]byte
	setCookiePath string
	contractAddr  common.Address
	contentDir    *string
	transactOpts  *bind.TransactOpts
	ctxCancel     context.CancelFunc

	callOpts = new(bind.CallOpts)
)

func main() {
	var (
		admin           = flag.String("admin", "", "admin e-mail address")
		listen          = flag.String("listen", ":80", "address for http requests")
		contractAddrStr = flag.String("contract", "", "deployed contract address (hex)")
		ethURL          = flag.String("ethurl", "http://localhost:8545/", "ethereum RPC URL")
		keyfile         = flag.String("keyfile", "", "key file")
		passphrase      = flag.String("passphrase", "", "passphrase")
	)
	contentDir = flag.String("content", ".", "directory with web assets")

	flag.Parse()

	if *admin != "" {
		// Send a link to the given e-mail address. Visiting that link
		// will set the user's admin cookie.

		_, err := rand.Read(adminSecret[:])
		if err != nil {
			log.Fatalf("computing admin secret: %s", err)
		}
		var setCookieSecret [16]byte
		_, err = rand.Read(setCookieSecret[:])
		if err != nil {
			log.Fatalf("computing cookie secret: %s", err)
		}
		setCookiePath = fmt.Sprintf("/setcookie/%x", setCookieSecret[:])
		err = sendmail(*admin, "Ninexd setcookie URL", fmt.Sprintf("http://%s%s", *listen, setCookiePath))
		if err != nil {
			log.Fatalf("sending setcookie e-mail: %s", err)
		}
	}

	keyReader, err := os.Open(*keyfile)
	if err != nil {
		log.Fatalf("opening keyfile: %s", err)
	}
	transactOpts, err = bind.NewTransactor(keyReader, *passphrase)
	if err != nil {
		log.Fatalf("creating transactor: %s", err)
	}
	keyReader.Close()

	transactOpts.GasLimit = big.NewInt(1000000)

	ethClient, err := ethclient.Dial(*ethURL)
	if err != nil {
		log.Fatalf("connecting to ethereum RPC service: %s", err)
	}

	contractAddrBytes, err := hex.DecodeString(*contractAddrStr)
	if err != nil {
		log.Fatalf("decoding contract address: %s", err)
	}

	copy(contractAddr[:], contractAddrBytes)

	nx, err = ninex.NewNinex(contractAddr, ethClient)
	if err != nil {
		log.Fatalf("instantiating contract: %s", err)
	}

	logCh := make(chan types.Log, 100)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	ctx := context.Background()
	ctx, ctxCancel = context.WithCancel(ctx)

	sub, err := ethClient.SubscribeFilterLogs(ctx, query, logCh)
	if err != nil {
		log.Fatalf("subscribing to log filter: %s", err)
	}

	go filter(ctx, logCh, sub.Err())

	pingCh := make(chan struct{}, 10)

	go serve(ctx, *listen, pingCh)
	go poll(ctx, pingCh)

	pingCh <- struct{}{}

	var (
		looping = true
		ticker  = time.Tick(time.Minute)
	)
	for looping {
		select {
		case <-ctx.Done():
			log.Printf("main loop exiting: %s", ctx.Err())
			looping = false

		case <-ticker:
			select {
			case pingCh <- struct{}{}:
			default:
			}
		}
	}
}

func sendmail(addr, subject, body string) error {
	// xxx todo
	log.Printf("sendmail %s [%s] %s", addr, subject, body)
	return nil
}
