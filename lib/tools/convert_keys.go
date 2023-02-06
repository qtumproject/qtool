package tools

import (
	"encoding/hex"
	"fmt"

	"github.com/qtumproject/qtool/lib/common"

	"github.com/btcsuite/btcutil/base58"
)

// ConvertPubkeyToAddrHash160 converts a compressed public key in hex format to a hash160 address
func ConvertPubkeyToAddrHash160(pubKeyString string) (*ConvertAddressResult, error) {
	prefix := pubKeyString[:2]
	pubKey, err := hex.DecodeString(pubKeyString)
	if err != nil {
		return nil, err
	}
	if prefix != "02" && prefix != "03" && prefix != "04" {
		return nil, fmt.Errorf("invalid public key prefix")
	}

	var compressedPubKey []byte

	if prefix == "02" || prefix == "03" {
		if len(pubKey) != 33 {
			return nil, fmt.Errorf("invalid public key length")
		} else {
			compressedPubKey = pubKey[:]
		}
	}
	if prefix == "04" {
		if len(pubKey) != 65 {
			return nil, fmt.Errorf("invalid public key length")
		} else {
			compressedPubKey = compressPublicKey(pubKey)
		}
	}
	hash160, err := common.HashPubKey(compressedPubKey)
	if err != nil {
		return nil, err
	}

	return &ConvertAddressResult{
		Address: hex.EncodeToString(hash160),
	}, nil
}

// ConvertPrivateKeyToWIF converts a private key in hex format to a base58 private key
func ConvertPrivateKeyToWIF(privKeyHex string, network string) (*ConvertPrivateKeyResult, error) {
	if len(privKeyHex) != 64 {
		return nil, fmt.Errorf("private key must be 64 characters long")
	}
	privKeyHexBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, err
	}

	version, err := getNetworkVersion(network)
	if err != nil {
		return nil, err
	}
	privKeyHexBytes = append([]byte{version}, privKeyHexBytes...)

	// Append the compression byte to the 33 byte private key
	privKeyHexBytes = append(privKeyHexBytes, byte(0x01))

	// Compute the checksum and append it to the 34 byte private key
	chksum := common.Checksum(privKeyHexBytes)
	privKeyHexBytes = append(privKeyHexBytes, chksum[:]...)

	privKeyB58 := base58.Encode(privKeyHexBytes)
	return &ConvertPrivateKeyResult{
		WIF: privKeyB58,
	}, nil
}

// ConvertPrivateKeyToHex converts a private key in base58 format to a hex private key
func ConvertPrivateKeyToHex(privKeyB58 string) (*ConvertPrivateKeyResult, error) {
	hexString, err := common.Base58ToHex(privKeyB58)
	if err != nil {
		return nil, err
	}

	if len(hexString) != 66 {
		return nil, fmt.Errorf("invalid private key length")
	}
	privKeyHex := hexString[:len(hexString)-2]
	return &ConvertPrivateKeyResult{
		PrivateKey: privKeyHex,
	}, nil
}
