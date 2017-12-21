package main

import "encoding/hex"

type hexBytes []byte

func (h hexBytes) MarshalText() ([]byte, error) {
	return []byte(hex.EncodeToString(h)), nil
}

func (h *hexBytes) UnmarshalText(inp []byte) error {
	n := hex.DecodedLen(len(inp))
	*h = make([]byte, n)
	_, err := hex.Decode(*h, inp)
	return err
}
