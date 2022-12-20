package keys

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/pkg/errors"
	"github.com/qtumproject/qtool/pkg/encoding"
	"github.com/qtumproject/qtool/pkg/tools"
)

const (
	// mainnet prefixes
	qtumMainPubKeyHashAddrID byte = 58
	qtumMainScriptHashAddrID byte = 50
	btcMainPubKeyHashAddrID  byte = 0x00
	mainnetVersion           byte = 0x80

	// testnet prefixes
	qtumTestNetPubKeyHashAddrID byte = 120 // hex: 0x78 dec: 120
	qtumTestNetScriptHashAddrID byte = 110 // hex: 0x6e dec: 110
	btcTestNetPubKeyHashAddrID  byte = 0x6F
	testnetVersion              byte = 0xEF
)

type KeyConverter struct {
	net string
}

func (c *KeyConverter) SetNetwork(net string) {
	c.net = net
}

// EncodeToBase58 encodes a private key hex string to a
// private key enconded as a base58 string
func (c *KeyConverter) EncodeToBase58(privKeyStr string) (string, error) {
	// Convert priveKey string to bytes
	privKey, err := encoding.HexStringToBytes32(privKeyStr)
	if err != nil {
		return "", errors.New("error decoding private key: " + err.Error())
	}

	// Append the network version byte to the 32 byte private key
	var version byte
	if c.net == "testnet" {
		version = testnetVersion
	} else if c.net == "mainnet" {
		version = mainnetVersion
	} else {
		return "", errors.New("invalid network")
	}
	privKey = append([]byte{version}, privKey...)

	// Append the compression byte to the 33 byte private key
	privKey = append(privKey, byte(0x01))

	// Compute the checksum and append it to the 34 byte private key
	chksum := tools.Checksum(privKey)
	privKey = append(privKey, chksum[:]...)
	privKeyBase58 := base58.Encode(privKey)
	return privKeyBase58, nil
}
