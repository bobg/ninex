package ninex

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type (
	Event interface {
		String() string
	}

	// This must track the definition of evGuess in ninex.sol
	EvGuess struct {
		Guesser    common.Address
		Commitment [32]byte
		Digits     []byte
		Stake      *big.Int
	}

	// This must track the definition of evWin in ninex.sol
	EvWin struct {
		Guesser    common.Address
		Value      *big.Int
		Digits     []byte
		Commitment [32]byte
	}

	// This must track the definition of evLose in ninex.sol
	EvLose struct {
		Guesser    common.Address
		Digits     []byte
		Commitment [32]byte
	}

	// This must track the definition of evWinByDefault in ninex.sol
	EvWinByDefault struct {
		Guesser common.Address
		Value   *big.Int
	}

	// This must track the definition of evNewCommitment in ninex.sol
	EvNewCommitment [32]byte

	// This must track the definition of evRevealed in ninex.sol
	EvRevealed struct {
		Commitment [32]byte
		Preimage   []byte
	}
)

func (ev *EvGuess) String() string {
	return fmt.Sprintf("guesser: %x; commitment: %x; digits: %s; stake: %s", ev.Guesser[:], ev.Commitment[:], string(ev.Digits), ev.Stake)
}

func (ev *EvWin) String() string {
	return fmt.Sprintf("guesser: %x; value %s; digits: %s; commitment: %x", ev.Guesser[:], ev.Value, string(ev.Digits), ev.Commitment[:])
}

func (ev *EvLose) String() string {
	return fmt.Sprintf("guesser: %x; digits: %s; commitment: %x", ev.Guesser[:], string(ev.Digits), ev.Commitment[:])
}

func (ev *EvWinByDefault) String() string {
	return fmt.Sprintf("guesser: %x; value: %s", ev.Guesser[:], ev.Value)
}

func (ev *EvNewCommitment) String() string {
	return fmt.Sprintf("commitment: %x", (*ev)[:])
}

func (ev *EvRevealed) String() string {
	return fmt.Sprintf("commitment: %x; preimage %s", ev.Commitment[:], string(ev.Preimage))
}
