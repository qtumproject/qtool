package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func MustMarshalIndent(v interface{}, prefix, indent string) []byte {
	res, err := json.MarshalIndent(v, prefix, indent)
	if err != nil {
		panic(err)
	}
	return res
}

type SampleValues struct {
	QtumAddressHex      string `json:"qtumAddressHex"`
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

// LoadSampleValues returns a struct with sample values for unit testing
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
