package tools

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	RAW_SCRIPT_PUBKEY_33 = "210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112ac"
	RAW_SCRIPT_PUBKEY_65 = "4104ae1a62fe09c5f51b13905f07f06b99a2f7159b2225f374cd378d71302fa28414e7aab37397f554a7df5f142c21c1b7303b8a0626f1baded5c72a704f7e6cd84cac"
	// scripPubKey = "210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112ac"
	ADDRESS_HEX_PUBKEY_65 = "0f41f27ebce08feb4da07e8d1888202746c4ca51"
	ADDRESS_B58_PUBKEY_65 = "qJx4J1AwTVowfUo6rxDiPvX3fykiUYZrbn"
)

func TestAddressFromP2PKScript(t *testing.T) {
	assert := assert.New(t)
	want := &P2pkToAddressResult{
		ScriptPubKey: ScriptPubKey{
			Hex: RAW_SCRIPT_PUBKEY_33,
			Asm: RAW_SCRIPT_PUBKEY_33[2:len(RAW_SCRIPT_PUBKEY_33)-2] + " OP_CHECKSIG",
		},
		PubKey:        RAW_SCRIPT_PUBKEY_33[2 : len(RAW_SCRIPT_PUBKEY_33)-2],
		AddressHex:    addressHex,
		AddressBase58: addressBase58,
	}
	t.Run("OP_DATA_33 valid input ", func(t *testing.T) {
		output, err := AddressFromP2PKScript(RAW_SCRIPT_PUBKEY_33, "qtum", "testnet")
		assert.NoError(err)
		assert.Equal(want, output)
	})
	t.Run("OP_DATA_33 invalid input length", func(t *testing.T) {
		output, err := AddressFromP2PKScript(RAW_SCRIPT_PUBKEY_33[1:], "qtum", "testnet")
		assert.Error(err)
		assert.Nil(output)
	})

	t.Run("OP_DATA_33 invalid input format", func(t *testing.T) {
		// pass a scriptPubKey with incorrect data length (65 instead of 33)
		invalidInput := strings.Replace(RAW_SCRIPT_PUBKEY_33, "21", "41", 1)
		output, err := AddressFromP2PKScript(invalidInput, "qtum", "testnet")
		assert.Error(err)
		assert.Nil(output)
	})
	//*******************************************************

	want = &P2pkToAddressResult{
		ScriptPubKey: ScriptPubKey{
			Hex: RAW_SCRIPT_PUBKEY_65,
			Asm: RAW_SCRIPT_PUBKEY_65[2:len(RAW_SCRIPT_PUBKEY_65)-2] + " OP_CHECKSIG",
		},
		PubKey:        RAW_SCRIPT_PUBKEY_65[2 : len(RAW_SCRIPT_PUBKEY_65)-2],
		AddressHex:    ADDRESS_HEX_PUBKEY_65,
		AddressBase58: ADDRESS_B58_PUBKEY_65,
	}
	t.Run("OP_DATA_65 valid input ", func(t *testing.T) {
		output, err := AddressFromP2PKScript(RAW_SCRIPT_PUBKEY_65, "qtum", "testnet")
		assert.NoError(err)
		assert.Equal(want, output)
	})
	t.Run("OP_DATA_65 invalid input length", func(t *testing.T) {
		output, err := AddressFromP2PKScript(RAW_SCRIPT_PUBKEY_65[1:], "qtum", "testnet")
		assert.Error(err)
		assert.Nil(output)
	})

	t.Run("OP_DATA_65 invalid input format", func(t *testing.T) {
		// pass a scriptPubKey with incorrect data length (65 instead of 33)
		invalidInput := strings.Replace(RAW_SCRIPT_PUBKEY_65, "41", "21", 1)
		output, err := AddressFromP2PKScript(invalidInput, "qtum", "testnet")
		assert.Error(err)
		assert.Nil(output)
	})
}
