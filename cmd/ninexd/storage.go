package main

import (
	"fmt"
	"log"
	"sync"
)

var (
	preimages [][]byte
	committed = -1
	storageMu sync.Mutex
)

func newPendingPreimage(preimage []byte) error {
	storageMu.Lock()
	defer storageMu.Unlock()

	preimages = append(preimages, preimage)
	log.Printf("new pending preimage %s", string(preimage))
	return nil
}

func commitPendingPreimage() error {
	storageMu.Lock()
	defer storageMu.Unlock()

	if committed >= len(preimages)-1 {
		return fmt.Errorf("committing pending preimage: committed is %d, len(preimages) is %d", committed, len(preimages))
	}
	committed++
	log.Printf("committed preimage %s", string(preimages[committed]))
	return nil
}

func getCommittedPreimage() ([]byte, error) {
	storageMu.Lock()
	defer storageMu.Unlock()

	if committed < 0 {
		return nil, fmt.Errorf("getting committed preimage: committed is %d", committed)
	}
	return preimages[committed], nil
}
