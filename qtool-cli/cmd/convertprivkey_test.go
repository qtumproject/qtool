package cmd

import (
	"testing"
)

func Test_ConvertPrivKey(t *testing.T) {
	s, err := loadSampleValues()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Convert Private Key to Hex", func(t *testing.T) {
		want := s.QtumPrivateKeyHex
		args := []string{"convertprivkey", s.QtumPrivateKeyWIF, "-f", "b58"}
		assertCmdSuccess(t, convertprivkeyCmd, args, want)

	})
	t.Run("Convert Private Key to Base58", func(t *testing.T) {
		want := s.QtumPrivateKeyWIF
		args := []string{"convertprivkey", s.QtumPrivateKeyHex, "-f", "hex", "-b", "qtum", "-n", "testnet"}
		assertCmdSuccess(t, convertprivkeyCmd, args, want)

	})
	t.Run("Convert Private Key to Base58 with bad input", func(t *testing.T) {
		args := []string{"convertprivkey", "abcd1234", "-f", "hex", "-b", "qtum", "-n", "testnet"}
		assertCmdError(t, convertprivkeyCmd, args)

	})
	t.Run("Convert Private Key to Hex with bad input", func(t *testing.T) {
		args := []string{"convertprivkey", "1234789", "-f", "b58", "-b", "qtum", "-n", "testnet"}
		assertCmdError(t, convertprivkeyCmd, args)

	})
}
