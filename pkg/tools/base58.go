package tools

import (
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
)

// TODO : delelte this file. Replaced by the one in the encoding package

// Base58ToHex converts a base58 string to a hex string
func Base58ToHex(input string) (string, error) {
	output, _, err := base58.CheckDecode(input)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(output), nil
}

// HexToBase58 converts a hex string to a base58 string
func HexToBase58(input string) (string, error) {
	output, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}

	return base58.Encode(output), nil
}
