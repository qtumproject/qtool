package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertAddressHexToBase58(t *testing.T) {
	assert := assert.New(t)
	t.Run("valid input ", func(t *testing.T) {
		output, err := ConvertAddressHexToBase58(addressHex, "qtum", "testnet")
		assert.NoError(err)
		want := &ConvertAddressResult{
			Address: addressBase58,
		}
		assert.Equal(want, output)
	})
	t.Run("invalid input length", func(t *testing.T) {
		output, err := ConvertAddressHexToBase58(addressHex[:len(addressHex)-2], "qtum", "testnet")
		assert.Error(err)
		assert.Nil(output)
	})
	t.Run("invalid input format", func(t *testing.T) {
		output, err := ConvertAddressHexToBase58("invalid", "qtum", "testnet")
		assert.Error(err)
		assert.Nil(output)
	})
	t.Run("invalid network", func(t *testing.T) {
		_, err := ConvertAddressHexToBase58(addressHex, "invalid", "invalid")
		assert.Error(err)
	})
	t.Run("wrong network", func(t *testing.T) {
		output, err := ConvertAddressHexToBase58(addressHex, "btc", "mainnet")
		assert.NoError(err)
		want := &ConvertAddressResult{
			Address: addressBase58,
		}
		assert.NotEqual(want, output)
	})
	t.Run("invalid address", func(t *testing.T) {
		_, err := ConvertAddressHexToBase58("invalid", "qtum", "testnet")
		assert.Error(err)
	})
}
func TestConvertAddressBase58ToHex(t *testing.T) {
	assert := assert.New(t)
	t.Run("valid input ", func(t *testing.T) {
		output, err := ConvertAddressBase58ToHex(addressBase58)
		assert.NoError(err)
		want := &ConvertAddressResult{
			Address: addressHex,
		}
		assert.Equal(want, output)
	})
	t.Run("invalid input length", func(t *testing.T) {
		output, err := ConvertAddressBase58ToHex(addressBase58[:len(addressBase58)-2])
		assert.Error(err)
		assert.Nil(output)
	})
	t.Run("invalid input format", func(t *testing.T) {
		output, err := ConvertAddressBase58ToHex("invalid")
		assert.Error(err)
		assert.Nil(output)
	})
}
