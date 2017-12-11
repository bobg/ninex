package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type info struct {
	commitment [32]byte
	bank       *big.Int

	commitmentSetTime int64
	firstGuessTime    int64
	lastGuessTime     int64
	revealedTime      int64

	afterCommitmentDelaySecs int64
	guessWindowSecs          int64
	afterLastGuessDelaySecs  int64
	afterRevealDelaySecs     int64
	revealTimeoutSecs        int64

	minGuessWei *big.Int
}

func getInfo() (*info, error) {
	info := new(info)

	var err error

	info.commitmentSetTime, err = secs(nx.MCommitmentSetTime)
	if err != nil {
		return nil, err
	}
	info.firstGuessTime, err = secs(nx.MFirstGuessTime)
	if err != nil {
		return nil, err
	}
	info.lastGuessTime, err = secs(nx.MLastGuessTime)
	if err != nil {
		return nil, err
	}
	info.revealedTime, err = secs(nx.MRevealedTime)
	if err != nil {
		return nil, err
	}

	info.afterCommitmentDelaySecs, err = secs(nx.MAfterCommitmentDelaySecs)
	if err != nil {
		return nil, err
	}
	info.guessWindowSecs, err = secs(nx.MGuessWindowSecs)
	if err != nil {
		return nil, err
	}
	info.afterLastGuessDelaySecs, err = secs(nx.MAfterLastGuessDelaySecs)
	if err != nil {
		return nil, err
	}
	info.afterRevealDelaySecs, err = secs(nx.MAfterRevealDelaySecs)
	if err != nil {
		return nil, err
	}
	info.revealTimeoutSecs, err = secs(nx.MRevealTimeoutSecs)
	if err != nil {
		return nil, err
	}

	info.commitment, err = nx.MCommitment(callOpts)
	if err != nil {
		return nil, err
	}
	info.bank, err = nx.MBank(callOpts)
	if err != nil {
		return nil, err
	}
	info.minGuessWei, err = nx.MMinGuessWei(callOpts)

	return info, err
}

func secs(f func(*bind.CallOpts) (*big.Int, error)) (int64, error) {
	secsBig, err := f(callOpts)
	if err != nil {
		return 0, err
	}
	if !secsBig.IsInt64() {
		return 0, fmt.Errorf("secs: %s does not fit in an int64", secsBig)
	}
	return secsBig.Int64(), nil
}
