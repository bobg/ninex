package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"log"
	"math/big"
	"time"
)

var bigTen = big.NewInt(10)

func poll(ctx context.Context, pingCh <-chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("poll exiting: %s", ctx.Err())
			return

		case <-pingCh:
			info, err := getInfo()
			if err != nil {
				log.Fatalf("getting current contract state: %s", err)
			}

			now := time.Now().Unix()

			if info.CommitmentSetTime == 0 || (info.RevealedTime > 0 && now >= (info.RevealedTime+info.AfterRevealDelaySecs)) {
				// set commitment

				var preimage [78]byte
				for i := 0; i < len(preimage); i++ {
					n, err := rand.Int(rand.Reader, bigTen)
					if err != nil {
						log.Fatalf("computing new preimage: %s", err)
					}
					preimage[i] = '0' + byte(n.Uint64())
				}
				err = newPendingPreimage(preimage[:])
				if err != nil {
					log.Fatalf("storing pending preimage: %s", err)
				}

				commitment := sha256.Sum256(preimage[:])
				_, err := nx.SetCommitment(transactOpts, commitment)
				if err != nil {
					log.Fatalf("sending setCommitment transaction: %s", err)
				}

				log.Printf("set commitment to %x", commitment[:])

				continue
			}

			if info.CommitmentSetTime > 0 && info.RevealedTime == 0 && info.LastGuessTime > 0 && now >= (info.LastGuessTime+info.AfterLastGuessDelaySecs) && now >= (info.FirstGuessTime+info.GuessWindowSecs) && now < (info.FirstGuessTime+info.RevealTimeoutSecs) {
				// reveal preimage

				preimage, err := getCommittedPreimage()
				if err != nil {
					log.Fatalf("getting committed preimage: %s", err)
				}

				_, err = nx.Reveal(transactOpts, preimage)
				if err != nil {
					log.Fatalf("sending reveal transaction: %s", err)
				}

				log.Printf("revealed preimage %x", preimage)

				continue
			}

			if info.CommitmentSetTime > 0 && info.RevealedTime == 0 && info.FirstGuessTime > 0 && now >= (info.FirstGuessTime+info.RevealTimeoutSecs) {
				// invoke a timeout

				_, err := nx.Timeout(transactOpts)
				if err != nil {
					log.Fatalf("sending timeout transaction: %s", err)
				}

				log.Print("contract commitment timed out")

				continue
			}
		}
	}
}
