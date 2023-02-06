package tools

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"
	"github.com/qtumproject/btcd/btcutil"
	"github.com/qtumproject/btcd/chaincfg"
	"github.com/qtumproject/qtool/lib/common"
)

// ConvertAddressHexToBase58() converts an address in hex format to a base58 address
//
// Note this conversion is dependent on the network (mainnet or testnet)
// and blockchain (bitcoin or qtum)
func ConvertAddressHexToBase58(addressHex string, blockchain, network string) (*ConvertAddressResult, error) {
	err := verifyBlockchain(blockchain)
	if err != nil {
		return nil, err
	}
	err = verifyNetwork(network)
	if err != nil {
		return nil, err
	}
	addressHex = common.RemoveHexPrefix(addressHex)
	if len(addressHex) != 40 {
		return nil, fmt.Errorf("ConvertAddressHexToBase58: addressHex must be 40 characters long")
	}
	addressBytes, err := hex.DecodeString(addressHex)
	if err != nil {
		return nil, err
	}

	addrID := getAddressID(blockchain, network)
	fullHex, err := appendChksumToPubKeyHash(addrID, addressBytes)
	if err != nil {
		return nil, err
	}

	return &ConvertAddressResult{
		Address: base58.Encode(fullHex),
	}, nil
}

// ConvertAddressHexToBase58() converts an address in hex format to a base58 address
//
// Note this conversion is dependent on the network (mainnet or testnet)
// and blockchain (bitcoin or qtum)
func AddressHexToBase58(addressHex string, chainParams *chaincfg.Params) (string, error) {
	err := verifyBlockchain(chainParams.Name)
	if err != nil {
		return "", err
	}
	err = verifyNetwork(chainParams.Net.String())
	if err != nil {
		return "", err
	}
	addressBytes, err := hex.DecodeString(strings.TrimPrefix(addressHex, "0x"))
	if err != nil {
		return "", errors.Wrapf(err, "Error decoding string to bytes address: %s", addressHex)
	}
	pubKeyHash, err := btcutil.NewAddressPubKeyHash(addressBytes, chainParams)
	if err != nil {
		return "", errors.Wrapf(err, "Error converting receiver address to base58: %s", addressHex)
	}
	return pubKeyHash.EncodeAddress(), nil

}

// ConvertAddressBase58ToHex converts an address in base58 format to a hex encoded address
func ConvertAddressBase58ToHex(addressBase58 string) (*ConvertAddressResult, error) {
	hexAddress, err := common.Base58ToHex(addressBase58)
	if err != nil {
		return nil, err
	}
	return &ConvertAddressResult{
		Address: hexAddress,
	}, nil
}
