package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type info struct {
	Commitment hexBytes `json:"commitment"`
	Preimage   string   `json:"preimage"`
	Bank       *big.Int `json:"bank"`

	CommitmentSetTime int64 `json:"commitment_set_time"`
	FirstGuessTime    int64 `json:"first_guess_time"`
	LastGuessTime     int64 `json:"last_guess_time"`
	RevealedTime      int64 `json:"revealed_time"`

	AfterCommitmentDelaySecs int64 `json:"after_commitment_delay_secs"`
	GuessWindowSecs          int64 `json:"guess_window_secs"`
	AfterLastGuessDelaySecs  int64 `json:"after_last_guess_delay_secs"`
	AfterRevealDelaySecs     int64 `json:"after_reveal_delay_secs"`
	RevealTimeoutSecs        int64 `json:"reveal_timeout_secs"`

	MinGuessWei *big.Int `json:"min_guess_wei"`
}

func getInfo() (*info, error) {
	info := new(info)

	var err error

	info.CommitmentSetTime, err = secs(nx.MCommitmentSetTime)
	if err != nil {
		return nil, err
	}
	info.FirstGuessTime, err = secs(nx.MFirstGuessTime)
	if err != nil {
		return nil, err
	}
	info.LastGuessTime, err = secs(nx.MLastGuessTime)
	if err != nil {
		return nil, err
	}
	info.RevealedTime, err = secs(nx.MRevealedTime)
	if err != nil {
		return nil, err
	}

	info.AfterCommitmentDelaySecs, err = secs(nx.MAfterCommitmentDelaySecs)
	if err != nil {
		return nil, err
	}
	info.GuessWindowSecs, err = secs(nx.MGuessWindowSecs)
	if err != nil {
		return nil, err
	}
	info.AfterLastGuessDelaySecs, err = secs(nx.MAfterLastGuessDelaySecs)
	if err != nil {
		return nil, err
	}
	info.AfterRevealDelaySecs, err = secs(nx.MAfterRevealDelaySecs)
	if err != nil {
		return nil, err
	}
	info.RevealTimeoutSecs, err = secs(nx.MRevealTimeoutSecs)
	if err != nil {
		return nil, err
	}

	commitment, err := nx.MCommitment(callOpts)
	if err != nil {
		return nil, err
	}
	info.Commitment = hexBytes(commitment[:])
	preimage, err := nx.MPreimage(callOpts)
	if err != nil {
		return nil, err
	}
	info.Preimage = string(preimage)
	info.Bank, err = nx.MBank(callOpts)
	if err != nil {
		return nil, err
	}
	info.MinGuessWei, err = nx.MMinGuessWei(callOpts)

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
