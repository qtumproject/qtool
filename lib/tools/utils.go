package tools

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/qtumproject/qtool/lib/common"
)

// appendChksumToPubKeyHash() takes the address ID and a 20 byte pubKeyHash
// and returns a byte array representing the full hex address
func appendChksumToPubKeyHash(qtumAddrID byte, pubKeyHash []byte) ([]byte, error) {
	if len(pubKeyHash) != 20 {
		return nil, errors.New("pubkey must be 20 bytes")
	}

	fullHex := append([]byte{qtumAddrID}, pubKeyHash...)

	chksum := common.Checksum(fullHex)
	fullHex = append(fullHex, chksum[:]...)

	return fullHex, nil
}

// verifyBlockchain() verifies that the blockchain is either "qtum" or "btc"
func verifyBlockchain(blockchain string) error {
	if blockchain != "qtum" && blockchain != "btc" {
		return errors.New("invalid blockchain: " + blockchain)
	}
	return nil
}

// verifyNetwork() verifies that the network is either "regtest", "testnet" or "mainnet"
func verifyNetwork(network string) error {

	if network != "testnet" && network != "mainnet" && network != "regtest" {
		return errors.New("invalid network: " + network)
	}
	return nil
}

// getAddressID() returns the corresponding address ID for
// the given blockchain and network
//
// Note: this is a helper function for ConvertAddressHexToBase58()
func getAddressID(blokchain, network string) byte {
	if network == "testnet" {
		if blokchain == "qtum" {
			return common.QtumTestNetPubKeyHashAddrID
		} else {
			return common.BtcTestNetPubKeyHashAddrID
		}
	} else {
		if blokchain == "qtum" {
			return common.QtumMainPubKeyHashAddrID
		} else {
			return common.BtcMainPubKeyHashAddrID
		}
	}
}

// getNetworkVersion() returns the corresponding network version for
// the given network
//
// Note: this is a helper function for ConvertPrivateKeyToWIF()
func getNetworkVersion(network string) (version byte, err error) {
	switch network {
	case "testnet":
		version = common.TestnetVersion
	case "mainnet":
		version = common.MainnetVersion
	case "regtest":
		version = common.TestnetVersion
	default:
		err = fmt.Errorf("invalid network: " + network)

	}
	return version, err
}
