package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"
	"ninex"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var commands = map[string]func([]string){
	"collect":        collect,
	"deploy":         deploy,
	"fund":           fund,
	"guess":          guess,
	"query":          query,
	"reveal":         reveal,
	"set-commitment": setCommitment,
	"set-params":     setParams,
	"timeout":        timeout,
	"withdraw":       withdraw,
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s command args", os.Args[0])
	}

	f, ok := commands[os.Args[1]]
	if !ok {
		log.Fatalf("unknown command %s", os.Args[1])
	}
	f(os.Args[2:])
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mustInt64(n *big.Int) int64 {
	if !n.IsInt64() {
		panic(fmt.Errorf("%s is not an int64", n))
	}
	return n.Int64()
}

func newFlagSet() (*flag.FlagSet, *string) {
	fs := new(flag.FlagSet)
	ethURL := fs.String("url", "http://localhost:8545/", "ethereum RPC URL")
	return fs, ethURL
}

func addKeyfilePassphrase(fs *flag.FlagSet) (keyfile, passphrase *string) {
	keyfile = fs.String("keyfile", "", "key file")
	passphrase = fs.String("passphrase", "", "passphrase")
	return keyfile, passphrase
}

func addContract(fs *flag.FlagSet) *string {
	return fs.String("contract", "", "deployed contract address (hex)")
}

func addValue(fs *flag.FlagSet) *uint64 {
	return fs.Uint64("value", 0, "value in wei")
}

func handleKeyfilePassphrase(keyfile, passphrase string) (*bind.TransactOpts, error) {
	keyReader, err := os.Open(keyfile)
	if err != nil {
		return nil, err
	}
	defer keyReader.Close()

	return bind.NewTransactor(keyReader, passphrase)
}

func handleContract(addrHex string, backend bind.ContractBackend) (*ninex.Ninex, error) {
	b, err := hex.DecodeString(addrHex)
	if err != nil {
		return nil, err
	}

	var contract common.Address
	copy(contract[:], b)

	return ninex.NewNinex(contract, backend)
}

func handleValue(wei uint64) *big.Int {
	val := new(big.Int)
	val.SetUint64(wei)
	return val
}
