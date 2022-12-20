package encoding

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

// HexStringToBytes32 converts a hex string to a byte array of length 32
// Input string must be 64 characters long
// Used to convert 256 bits keys and addresses
func HexStringToBytes32(dataStr string) ([]byte, error) {
	if len(dataStr) != 64 {
		return nil, fmt.Errorf("data must be 64 characters long")
	}
	dataByte, err := hex.DecodeString(dataStr)
	if err != nil {
		return nil, err
	}
	return dataByte, nil

}

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
