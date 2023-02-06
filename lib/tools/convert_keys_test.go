package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertPubkeyToAddrHash160(t *testing.T) {
	assert := assert.New(t)
	t.Run("valid input ", func(t *testing.T) {
		output, err := ConvertPubkeyToAddrHash160(pubKey)
		assert.NoError(err)
		want := &ConvertAddressResult{
			Address: addressHex,
		}
		assert.Equal(want, output)
	})

	t.Run("invalid input length", func(t *testing.T) {
		output, err := ConvertPubkeyToAddrHash160("invalid")
		assert.Error(err)
		assert.Nil(output)
	})

	t.Run("invalid input format", func(t *testing.T) {
		output, err := ConvertPubkeyToAddrHash160("invalid")
		assert.Error(err)
		assert.Nil(output)
	})
}

func TestConvertPrivateKeyToWIF(t *testing.T) {
	assert := assert.New(t)
	t.Run("valid input ", func(t *testing.T) {
		output, err := ConvertPrivateKeyToWIF(privateKeyHex, "testnet")
		assert.NoError(err)
		want := &ConvertPrivateKeyResult{
			PrivateKey: privateKeyWIF,
		}
		assert.Equal(want, output)
	})
	t.Run("invalid input length", func(t *testing.T) {
		output, err := ConvertPrivateKeyToWIF(privateKeyHex[:len(privateKeyHex)-2], "testnet")
		assert.Error(err)
		assert.Nil(output)
	})
	t.Run("invalid input format", func(t *testing.T) {
		output, err := ConvertPrivateKeyToWIF("invalid", "testnet")
		assert.Error(err)
		assert.Nil(output)
	})
	t.Run("invalid network", func(t *testing.T) {
		_, err := ConvertPrivateKeyToWIF(privateKeyHex, "invalid")
		assert.Error(err)
	})
	t.Run("wrong network", func(t *testing.T) {
		output, err := ConvertPrivateKeyToWIF(privateKeyHex, "mainnet")
		assert.NoError(err)
		want := &ConvertPrivateKeyResult{
			PrivateKey: privateKeyWIF,
		}
		assert.NotEqual(want, output)
	})
}

func TestConvertPrivateKeyToHex(t *testing.T) {
	assert := assert.New(t)
	t.Run("valid input ", func(t *testing.T) {
		output, err := ConvertPrivateKeyToHex(privateKeyWIF)
		assert.NoError(err)
		want := &ConvertPrivateKeyResult{
			PrivateKey: privateKeyHex,
		}
		assert.Equal(want, output)
	})
	t.Run("invalid input length", func(t *testing.T) {
		output, err := ConvertPrivateKeyToHex(privateKeyWIF[:len(privateKeyWIF)-2])
		assert.Error(err)
		assert.Nil(output)
	})
	t.Run("invalid input format", func(t *testing.T) {
		output, err := ConvertPrivateKeyToHex("invalid")
		assert.Error(err)
		assert.Nil(output)
	})
}

func TestGetAddressFromPrivKey(t *testing.T) {
	testHelperGetAddressFromPrivKey(t, "hex", privateKeyHex, addressBase58)
	testHelperGetAddressFromPrivKey(t, "b58", privateKeyWIF, addressBase58)

}

func testHelperGetAddressFromPrivKey(t *testing.T, format string, privateKey string, want string) {
	t.Helper()
	assert := assert.New(t)
	t.Run("valid input ", func(t *testing.T) {
		output, err := GetAddressFromPrivkey(
			privateKey,
			"qtum",
			"testnet",
			format,
		)
		assert.NoError(err)
		if assert.NotNil(output) {
			assert.Equal(output.AddressBase58, addressBase58)
		}
	})
	t.Run("invalid input ", func(t *testing.T) {
		output, err := GetAddressFromPrivkey(
			privateKey[:len(privateKey)-2],
			"qtum",
			"testnet",
			format,
		)
		assert.Error(err)
		assert.Nil(output)
	})
	t.Run("invalid format ", func(t *testing.T) {
		output, err := GetAddressFromPrivkey(
			privateKey,
			"qtum",
			"testnet",
			"invalid",
		)
		assert.Error(err)
		assert.Nil(output)
	})

	var wrongFormat string
	if format == "hex" {
		wrongFormat = "b58"
	} else {
		wrongFormat = "hex"
	}
	t.Run("wrong format ", func(t *testing.T) {
		output, err := GetAddressFromPrivkey(
			privateKey,
			"qtum",
			"testnet",
			wrongFormat,
		)
		assert.Error(err)
		assert.Nil(output)
	})
}
