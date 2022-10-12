package tools

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
)

// ConvertAddressHexToBase58 converts an address in hex format to a base58 address
// Note this conversion is dependent on the network (mainnet or testnet) and blockchain (bitcoin or qtum)
func ConvertAddressHexToBase58(addressHex string, blockchain, network string) (*ConvertAddressResult, error) {
	err := verifyBlockchain(blockchain)
	if err != nil {
		return nil, err
	}
	err = verifyNetwork(network)
	if err != nil {
		return nil, err
	}
	if len(addressHex) != 40 {
		return nil, fmt.Errorf("ConvertAddressHexToBase58: addressHex must be 40 characters long")
	}
	addressBytes, err := hex.DecodeString(addressHex)
	if err != nil {
		return nil, err
	}

	addrID := getAddressID(blockchain, network)
	fullHex := appendChksumToAddrHash160(addrID, addressBytes)

	return &ConvertAddressResult{
		Address: base58.Encode(fullHex),
	}, nil
}

// ConvertAddressBase58ToHex converts an address in base58 format to a hex encoded address
func ConvertAddressBase58ToHex(addressBase58 string) (*ConvertAddressResult, error) {
	hexAddress, err := Base58ToHex(addressBase58)
	if err != nil {
		return nil, err
	}
	return &ConvertAddressResult{
		Address: hexAddress,
	}, nil
}
