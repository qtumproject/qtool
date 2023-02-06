// Package common contains various helper functions.
package common

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

// Converts an int to a hex string (with 0x prefix).
func EncodeBig(input *big.Int) string {
	return hexutil.EncodeBig(input)
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

// HexStringToBytes32 converts a hex string to a byte array of length 32.
//
// If input string is not 64 characters long returns an error.
//
// Used to convert 256 bits keys and addresses to byte array.
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

// IsValidHexAddress checks if the given address is a valid ethereum hex address.
func IsValidHexAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}
