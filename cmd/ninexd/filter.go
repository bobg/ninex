package main

import (
	"bytes"
	"log"
	"ninex"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

func filter(logCh <-chan types.Log, errCh <-chan error) {
	abi, err := abi.JSON(bytes.NewReader([]byte(ninex.NinexABI)))
	if err != nil {
		log.Fatalf("instantiating ABI: %s", err)
	}

	for {
		select {
		case l, ok := <-logCh:
			if !ok {
				log.Fatal("subscription channel closed")
			}
		Outer:
			for _, t := range l.Topics {
				for name, e := range abi.Events {
					if t == e.Id() {
						var ev ninex.Event
						switch e.Name {
						case "evGuess":
							ev = new(ninex.EvGuess)
						case "evWin":
							ev = new(ninex.EvWin)
						case "evLose":
							ev = new(ninex.EvLose)
						case "evWinByDefault":
							ev = new(ninex.EvWinByDefault)
						case "evNewCommitment":
							ev = new(ninex.EvNewCommitment)
							err := commitPendingPreimage()
							if err != nil {
								log.Fatalf("committing pending preimage: %s", err)
							}
						case "evRevealed":
							ev = new(ninex.EvRevealed)
						default:
							id := e.Id()
							log.Fatalf("unknown event %s, id %x, tx %x", name, id[:], l.TxHash[:])
						}
						err = abi.Unpack(ev, name, l.Data)
						if err != nil {
							id := e.Id()
							log.Printf("failed to unpack %s event, id %x, tx %x, data %x", name, id[:], l.TxHash[:], l.Data)
						} else {
							log.Printf("tx %x event %s [%s]", l.TxHash[:], name, ev)
						}
						continue Outer
					}
				}
			}

		case err, ok := <-errCh:
			if ok {
				log.Fatalf("subscription error: %s", err)
			}
		}
	}
}
