// Package common contains various helper functions.
package common

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/common/hexutil"
)

func RemoveHexPrefix(hex string) string {
	if strings.HasPrefix(hex, "0x") {
		return hex[2:]
	}
	return hex
}
func AddHexPrefix(hex string) string {
	if strings.HasPrefix(hex, "0x") {
		return hex
	}
	return "0x" + hex
}

// Decodes a hex string (with or w/o 0x prefix) into an int .
func DecodeBig(input string) (*big.Int, error) {
	input = AddHexPrefix(input)
	return hexutil.DecodeBig(input)
}

// Converts a hex string (without 0x prefix) to a hex string in Big Endian format
func ConvertToBigEndian(hex string) (string, error) {
	if len(hex)%2 != 0 {
		return "", fmt.Errorf("invalid hex string")
	}
	var result string
	for i := len(hex); i > 0; i -= 2 {
		result += hex[i-2 : i]
	}
	// trim leading zero
	return strings.TrimLeft(result, "0"), nil
}
