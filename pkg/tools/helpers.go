package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// appendChksumToAddrHash160 takes a hash160 address and returns a hex array that contains the prefix and checksum
func appendChksumToAddrHash160(qtumPubKeyAddrID byte, hash160 []byte) []byte {

	IDHash160 := append([]byte{qtumPubKeyAddrID}, hash160...)

	chksum := Checksum(IDHash160)
	addressFullHex := append(IDHash160, chksum[:]...)

	return addressFullHex
}

func verifyBlockchain(blockchain string) error {
	if blockchain != "qtum" && blockchain != "btc" {
		return fmt.Errorf("invalid blockchain")
	}
	return nil
}

func verifyNetwork(network string) error {

	if network != "testnet" && network != "mainnet" {
		return fmt.Errorf("invalid network")
	}
	return nil
}

func getAddressID(blokchain, network string) byte {
	if network == "testnet" {
		if blokchain == "qtum" {
			return qtumTestNetPubKeyHashAddrID
		} else {
			return btcTestNetPubKeyHashAddrID
		}
	} else {
		if blokchain == "qtum" {
			return qtumMainPubKeyHashAddrID
		} else {
			return btcMainPubKeyHashAddrID
		}
	}
}

type SampleValues struct {
	QtumAddressHex      string `json:"qtumAddressHex"`    // "qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW"
	QtumAddressBase58   string `json:"qtumAddressBase58"` // "qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW"
	QtumPrivateKeyWIF   string `json:"qtumPrivateKeyWIF"`
	QtumPrivateKeyHex   string `json:"qtumPrivateKeyHex"`
	QtumPubKey          string `json:"qtumPubKey"`
	QtumScriptpubkey_33 string `json:"qtumScriptpubkey_33"`
	QtumScriptpubkey_65 string `json:"qtumScriptpubkey_65"`
	BtcAddressHex       string `json:"btcAddressHex"`
	BtcAddressBase58    string `json:"btcAddressBase58"`
	BtcPrivateKeyWIF    string `json:"btcPrivateKeyWIF"`
	BtcPrivateKeyHex    string `json:"btcPrivateKeyHex"`
	BtcPubKey           string `json:"btcPubKey"`
}

// LoadSampleValues returns a struct with sample values for testing
func LoadSampleValues(filename string) (*SampleValues, error) {
	samplesFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer samplesFile.Close()

	samplesBytes, err := ioutil.ReadAll(samplesFile)
	if err != nil {
		return nil, err
	}

	var samples SampleValues

	err = json.Unmarshal([]byte(samplesBytes), &samples)

	if err != nil {
		return nil, err
	}
	return &samples, nil
}
