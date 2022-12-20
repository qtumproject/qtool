package tools

import (
	"encoding/hex"
	"fmt"

	"crypto/sha256"

	"github.com/qtumproject/btcd/btcec/v2"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
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
	hash1 := sha256.Sum256(compressedPubKey)
	hasher160 := ripemd160.New()
	_, err = hasher160.Write(hash1[:])
	if err != nil {
		return nil, err
	}
	hash2 := hasher160.Sum(nil)

	return &ConvertAddressResult{
		Address: hex.EncodeToString(hash2),
	}, nil
}

// ConvertPrivateKeyToWIF converts a private key in hex (SHA256) format to a base58 private key
func ConvertPrivateKeyToWIF(privKeyHex string, network string) (*ConvertPrivateKeyResult, error) {
	if len(privKeyHex) != 64 {
		return nil, fmt.Errorf("private key must be 64 characters long")
	}
	privKeyHexBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, err
	}

	// Append the network version byte to the 32 byte private key
	var version byte
	if network == "testnet" {
		version = testnetVersion
	} else if network == "mainnet" {
		version = mainnetVersion
	} else {
		return nil, fmt.Errorf("invalid network")
	}
	privKeyHexBytes = append([]byte{version}, privKeyHexBytes...)

	// Append the compression byte to the 33 byte private key
	privKeyHexBytes = append(privKeyHexBytes, byte(0x01))

	// Compute the checksum and append it to the 34 byte private key
	chksum := Checksum(privKeyHexBytes)
	privKeyHexBytes = append(privKeyHexBytes, chksum[:]...)

	privKeyB58 := base58.Encode(privKeyHexBytes)
	return &ConvertPrivateKeyResult{
		PrivateKey: privKeyB58,
	}, nil
}

// ConvertPrivateKeyToHex converts a private key in base58 format to a hex private key
func ConvertPrivateKeyToHex(privKeyB58 string) (*ConvertPrivateKeyResult, error) {
	hexString, err := Base58ToHex(privKeyB58)
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

// GetAddressFromPrivkey gets the corresponding base58 (public) address from the private key
func GetAddressFromPrivkey(
	privKeyStr string,
	blockchain string,
	network string,
	privKeyFormat string) (*GetAddressFromPrivkeyResult, error) {
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
