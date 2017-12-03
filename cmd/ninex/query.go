package main

import (
	"log"
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

	escrow, err := nx.MEscrow(callOpts)
	must(err)
	log.Printf("escrow: %s", escrow)

	guessedBy, err := nx.MGuessedBy(callOpts)
	must(err)
	log.Printf("guessed by: %x", guessedBy[:])

	guessedDigits, err := nx.MGuessedDigits(callOpts)
	must(err)
	log.Printf("guessed digits: %s", string(guessedDigits))

	commitmentSetTime, err := nx.MCommitmentSetTime(callOpts)
	must(err)
	if commitmentSetTime.Sign() == 0 {
		log.Print("commitment-set time: <unset>")
	} else {
		log.Printf("commitment-set time: %s", time.Unix(mustInt64(commitmentSetTime), 0))
	}

	guessedTime, err := nx.MGuessedTime(callOpts)
	must(err)
	if guessedTime.Sign() == 0 {
		log.Print("guessed time: <unset>")
	} else {
		log.Printf("guessed time: %s", time.Unix(mustInt64(guessedTime), 0))
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

	afterGuessDelaySecs, err := nx.MAfterGuessDelaySecs(callOpts)
	must(err)
	log.Printf("after guess delay: %s secs", afterGuessDelaySecs)

	afterRevealDelaySecs, err := nx.MAfterRevealDelaySecs(callOpts)
	must(err)
	log.Printf("after reveal delay: %s secs", afterRevealDelaySecs)

	revealTimeoutSecs, err := nx.MRevealTimeoutSecs(callOpts)
	must(err)
	log.Printf("reveal timeout: %s secs", revealTimeoutSecs)

	minGuess, err := nx.MMinGuess(callOpts)
	must(err)
	log.Printf("min guess: %s", minGuess)
}
