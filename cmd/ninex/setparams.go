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
		guessWindowSecs          = fs.Uint("guess-window-secs", 0, "guess window")
		afterLastGuessDelaySecs  = fs.Uint("after-last-guess-delay-secs", 0, "after-last-guess delay")
		afterRevealDelaySecs     = fs.Uint("after-reveal-delay-secs", 0, "after-reveal delay")
		revealTimeoutSecs        = fs.Uint("reveal-timeout-secs", 0, "reveal timeout")
		minGuessWeiStr           = fs.String("min-guess-wei", "", "min guess in wei")
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

	if *guessWindowSecs > 0 {
		tx, err := nx.SetGuessWindow(transactOpts, asBigInt(guessWindowSecs))
		must(err)

		receipt, err := bind.WaitMined(ctx, client, tx)
		must(err)
		log.Printf("setting mGuessWindowSecs: %s", receipt)
	}

	if *afterLastGuessDelaySecs > 0 {
		tx, err := nx.SetAfterLastGuessDelay(transactOpts, asBigInt(afterLastGuessDelaySecs))
		must(err)

		receipt, err := bind.WaitMined(ctx, client, tx)
		must(err)
		log.Printf("setting mAfterLastGuessDelaySecs: %s", receipt)
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

	if *minGuessWeiStr != "" {
		var minGuessWei big.Int
		_, ok := minGuessWei.SetString(*minGuessWeiStr, 10)
		if !ok || minGuessWei.Sign() < 1 {
			panic(fmt.Errorf("%s is not a positive base 10 number", *minGuessWeiStr))
		}

		tx, err := nx.SetMinGuessWei(transactOpts, &minGuessWei)
		must(err)

		receipt, err := bind.WaitMined(ctx, client, tx)
		must(err)
		log.Printf("setting mMinGuessWei: %s", receipt)
	}
}

func asBigInt(n *uint) *big.Int {
	res := new(big.Int)
	res.SetUint64(uint64(*n))
	return res
}
