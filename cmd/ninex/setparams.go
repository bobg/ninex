package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func setParams(args []string) {
	fs, ethURL := newFlagSet()

	var (
		keyfile, passphrase      = addKeyfilePassphrase(fs)
		contractAddr             = addContract(fs)
		afterCommitmentDelaySecs = fs.Uint("after-commitment-delay-secs", 0, "after-commitment delay")
		afterGuessDelaySecs      = fs.Uint("after-guess-delay-secs", 0, "after-guess delay")
		afterRevealDelaySecs     = fs.Uint("after-reveal-delay-secs", 0, "after-reveal delay")
		revealTimeoutSecs        = fs.Uint("reveal-timeout-secs", 0, "reveal timeout")
		minGuessStr              = fs.String("min-guess", "", "min guess")
	)

	err := fs.Parse(args)
	must(err)

	transactOpts, err := handleKeyfilePassphrase(*keyfile, *passphrase)
	must(err)

	client, err := ethclient.Dial(*ethURL)
	must(err)

	nx, err := handleContract(*contractAddr, client)
	must(err)

	ctx := context.Background()

	if *afterCommitmentDelaySecs > 0 {
		tx, err := nx.SetAfterCommitmentDelay(transactOpts, asBigInt(afterCommitmentDelaySecs))
		must(err)

		receipt, err := bind.WaitMined(ctx, client, tx)
		must(err)
		log.Printf("setting mAfterCommitmentDelaySecs: %s", receipt)
	}

	if *afterGuessDelaySecs > 0 {
		tx, err := nx.SetAfterGuessDelay(transactOpts, asBigInt(afterGuessDelaySecs))
		must(err)

		receipt, err := bind.WaitMined(ctx, client, tx)
		must(err)
		log.Printf("setting mAfterGuessDelaySecs: %s", receipt)
	}

	if *afterRevealDelaySecs > 0 {
		tx, err := nx.SetAfterRevealDelay(transactOpts, asBigInt(afterRevealDelaySecs))
		must(err)

		receipt, err := bind.WaitMined(ctx, client, tx)
		must(err)
		log.Printf("setting mAfterRevealDelaySecs: %s", receipt)
	}

	if *revealTimeoutSecs > 0 {
		tx, err := nx.SetRevealTimeout(transactOpts, asBigInt(revealTimeoutSecs))
		must(err)

		receipt, err := bind.WaitMined(ctx, client, tx)
		must(err)
		log.Printf("setting mRevealTimeoutSecs: %s", receipt)
	}

	if *minGuessStr != "" {
		var minGuess big.Int
		_, ok := minGuess.SetString(*minGuessStr, 10)
		if !ok || minGuess.Sign() < 1 {
			panic(fmt.Errorf("%s is not a positive base 10 number", *minGuessStr))
		}

		tx, err := nx.SetMinGuess(transactOpts, &minGuess)
		must(err)

		receipt, err := bind.WaitMined(ctx, client, tx)
		must(err)
		log.Printf("setting mMinGuess: %s", receipt)
	}
}

func asBigInt(n *uint) *big.Int {
	res := new(big.Int)
	res.SetUint64(uint64(*n))
	return res
}
