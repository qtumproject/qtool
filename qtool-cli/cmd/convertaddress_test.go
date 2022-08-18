package cmd

import (
	"testing"
)

func Test_ConvertAddress(t *testing.T) {
	s, err := loadSampleValues()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Convert Address to Hex", func(t *testing.T) {
		want := s.QtumAddressHex
		args := []string{"convertaddress", s.QtumAddressBase58, "-f", "b58"}
		assertCmdSuccess(t, convertaddressCmd, args, want)

	})
	t.Run("Convert Address to Base58", func(t *testing.T) {
		want := s.QtumAddressBase58
		args := []string{"convertaddress", s.QtumAddressHex, "-f", "hex", "-b", "qtum", "-n", "testnet"}
		assertCmdSuccess(t, convertaddressCmd, args, want)

	})
	t.Run("Convert Address to Base58 with bad input", func(t *testing.T) {
		args := []string{"convertaddress", "abcd1234", "-f", "hex", "-b", "qtum", "-n", "testnet"}
		assertCmdError(t, convertaddressCmd, args)

	})
	t.Run("Convert Address to Hex with bad input", func(t *testing.T) {
		args := []string{"convertaddress", "1234789", "-f", "b58", "-b", "qtum", "-n", "testnet"}
		assertCmdError(t, convertaddressCmd, args)

	})
}
