package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func query(args []string) {
	fs, ethURL := newFlagSet()
	contractAddr := addContract(fs)

	err := fs.Parse(args)
	must(err)

	client, err := ethclient.Dial(*ethURL)
	must(err)

	nx, err := handleContract(*contractAddr, client)
	must(err)

	callOpts := new(bind.CallOpts)

	commitment, err := nx.MCommitment(callOpts)
	must(err)
	log.Printf("commitment: %x", commitment[:])

	bank, err := nx.MBank(callOpts)
	must(err)
	log.Printf("bank: %s", bank)

	numGuessesBig, err := nx.NumGuesses(callOpts)
	must(err)
	if !numGuessesBig.IsInt64() {
		panic(fmt.Errorf("num guesses %s is not an int64", numGuessesBig))
	}
	numGuesses := numGuessesBig.Int64()
	if numGuesses > 0 {
		log.Print("guesses:")
		for i := int64(0); i < numGuesses; i++ {
			guess, err := nx.MGuesses(callOpts, big.NewInt(i))
			must(err)
			log.Printf("  {guesser: %x", guess.Guesser[:])
			log.Printf("   digits: %s", string(guess.Digits))
			log.Printf("   escrow: %s}", guess.Escrow)
		}
	}

	commitmentSetTime, err := nx.MCommitmentSetTime(callOpts)
	must(err)
	if commitmentSetTime.Sign() == 0 {
		log.Print("commitment-set time: <unset>")
	} else {
		log.Printf("commitment-set time: %s", time.Unix(mustInt64(commitmentSetTime), 0))
	}

	firstGuessTime, err := nx.MFirstGuessTime(callOpts)
	must(err)
	if firstGuessTime.Sign() == 0 {
		log.Print("first-guess time: <unset>")
	} else {
		log.Printf("first-guess time: %s", time.Unix(mustInt64(firstGuessTime), 0))
	}

	lastGuessTime, err := nx.MLastGuessTime(callOpts)
	must(err)
	if lastGuessTime.Sign() == 0 {
		log.Print("last-guess time: <unset>")
	} else {
		log.Printf("last-guess time: %s", time.Unix(mustInt64(lastGuessTime), 0))
	}

	revealedTime, err := nx.MRevealedTime(callOpts)
	must(err)
	if revealedTime.Sign() == 0 {
		log.Print("revealed time: <unset>")
	} else {
		log.Printf("revealed time: %s", time.Unix(mustInt64(revealedTime), 0))
	}

	afterCommitmentDelaySecs, err := nx.MAfterCommitmentDelaySecs(callOpts)
	must(err)
	log.Printf("after commitment delay: %s secs", afterCommitmentDelaySecs)

	guessWindowSecs, err := nx.MGuessWindowSecs(callOpts)
	must(err)
	log.Printf("guess window: %s secs", guessWindowSecs)

	afterLastGuessDelaySecs, err := nx.MAfterLastGuessDelaySecs(callOpts)
	must(err)
	log.Printf("after last-guess delay: %s secs", afterLastGuessDelaySecs)

	afterRevealDelaySecs, err := nx.MAfterRevealDelaySecs(callOpts)
	must(err)
	log.Printf("after reveal delay: %s secs", afterRevealDelaySecs)

	revealTimeoutSecs, err := nx.MRevealTimeoutSecs(callOpts)
	must(err)
	log.Printf("reveal timeout: %s secs", revealTimeoutSecs)

	minGuessWei, err := nx.MMinGuessWei(callOpts)
	must(err)
	log.Printf("min guess: %s wei", minGuessWei)
}
