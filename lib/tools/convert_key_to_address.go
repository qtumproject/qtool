package tools

import (
	"encoding/hex"
	"fmt"

	"github.com/qtumproject/btcd/btcec/v2"
)

// GetAddressFromPrivkey derives the corresponding base58 address from the private key
//
// Parameters:
//
// privKeyStr: the private key in hex or base58 format
//
// blockchain: the blockchain to use (qtum, bitcoin)
//
// network: the network to use (mainnet, testnet, regtest)
//
// privKeyFormat: the format of the private key ('hex' or 'b58')
func GetAddressFromPrivkey(
	privKeyStr string,
	blockchain string,
	network string,
	privKeyFormat string) (*GetAddressFromPrivkeyResult, error) {

	err := verifyBlockchain(blockchain)
	if err != nil {
		return nil, err
	}
	err = verifyNetwork(network)
	if err != nil {
		return nil, err
	}

	var privateKeyHex string
	var privateKeyWIF string

	if privKeyFormat == "b58" {
		privateKeyWIF = privKeyStr
		result, err := ConvertPrivateKeyToHex(privKeyStr)
		if err != nil {
			return nil, err
		}
		privateKeyHex = result.PrivateKey
	} else if privKeyFormat == "hex" {
		privateKeyHex = privKeyStr
		result, err := ConvertPrivateKeyToWIF(privKeyStr, network)
		if err != nil {
			return nil, err
		}
		privateKeyWIF = result.PrivateKey
	} else {
		return nil, fmt.Errorf("invalid private key format")
	}

	ecPrivKey, err := ecdsaPrivKeyFromString(privateKeyHex)
	if err != nil {
		return nil, err
	}
	ecPubKey := ecPrivKey.PubKey()
	ecPubKeyCompressed := ecPubKey.SerializeCompressed()
	ecPubKeyCompressedStr := hex.EncodeToString(ecPubKeyCompressed)

	pubKeyToHexAddrResult, err := ConvertPubkeyToAddrHash160(ecPubKeyCompressedStr)
	if err != nil {
		return nil, err
	}

	addrHexToB58Result, err := ConvertAddressHexToBase58(pubKeyToHexAddrResult.Address, blockchain, network)
	if err != nil {
		return nil, err
	}
	return &GetAddressFromPrivkeyResult{
		PrivateKeyHex: privateKeyHex,
		PrivateKeyWIF: privateKeyWIF,
		PublicKeyHex:  ecPubKeyCompressedStr,
		AddressHex:    pubKeyToHexAddrResult.Address,
		AddressBase58: addrHexToB58Result.Address,
	}, nil
}

// ecdsaPrivKeyFromString converts a private key in hex format to a btcec.PrivateKey
func ecdsaPrivKeyFromString(privKeyHexStr string) (*btcec.PrivateKey, error) {
	pkBytes, err := hex.DecodeString(privKeyHexStr)
	if err != nil {
		return nil, err
	}
	privKey, _ := btcec.PrivKeyFromBytes(pkBytes)

	return privKey, nil
}

func compressPublicKey(pubkey []byte) []byte {
	var prefix byte
	if pubkey[64]&0x1 == 1 {
		prefix = 0x03
	} else {
		prefix = 0x02
	}
	return append([]byte{prefix}, pubkey[1:33]...)
}
