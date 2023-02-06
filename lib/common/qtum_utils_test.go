package common

import (
	"math/big"
	"strings"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

type Tests struct {
	Name    string
	Wei     string
	Dec     string
	Satoshi int
	Qtum    float64
}

func TestConversions(t *testing.T) {
	assert := assert.New(t)
	/*
		⇒ 1 satoshi == 100.000.000 wei
		⇒ 1000 satoshi == 10.000.000.000 wei
		⇒ 1 BTC == 100.000.000 satoshi == 1.000.000.000.000.000.000 wei
		⇒ 0.1 BTC == 10.000.000 satoshi == 100.000.000.000.000.000 wei
		⇒ 1 x10^8 wei = 100.000.000 wei  == 1 satoshi == 0.00000001 BTC
		⇒ 1 x10^9 wei = 1000.000.000 wei  == 10 satoshi == 0.0000001 BTC
		⇒ 1 x10^15 wei = 1.0000.000.000.000.000 wei  == 10.000.000 satoshi == 0.1 BTC
		⇒ 1 x10^16 wei = 10.0000.000.000.000.000 wei  == 100.000.000 satoshi == 1 BTC
		⇒ 1 x10^18 wei = 1000.0000.000.000.000.000 wei  == 10.000.000.000 satoshi == 100 BTC
	*/

	var tests = []Tests{
		{"0 wei", "0x0", "0", 0, 0},
		{"1 wei", "0x1", "1", 0, 0},
		{"1 wei x 10^8", "0x5F5E100", "100000000", 0, 0},
		{"1 wei x 10^9", "0x3B9ACA00", "1000000000", 0, 0},
		{"1 wei x 10^10", "0x2540BE400", "10000000000", 1, 0.00000001},
		{"1 wei x 10^15", "0x38D7EA4C68000", "1000000000000000", 100000, 0.001},
		{"1 wei x 10^16", "0x2386F26FC10000", "10000000000000000", 1000000, 0.01},
		{"1 wei x 10^18", "0xDE0B6B3A7640000", "1000000000000000000", 100000000, 1},
	}

	for _, test := range tests {
		t.Run(test.Name+": Convert Hex From Wei To Satoshi", func(t *testing.T) {
			inSat, err := ConvertHexFromWeiToSatoshi(test.Wei)
			assert.NoError(err)
			satoshis, err := DecodeBig(inSat)
			assert.NoError(err)
			assert.Equal(test.Satoshi, int(satoshis.Int64()))
		})
		t.Run(test.Name+": Convert From Satoshi To Qtum", func(t *testing.T) {
			inSat := decimal.NewFromBigInt(big.NewInt(int64(test.Satoshi)), 0)
			inQtum := ConvertFromSatoshisToQtum(inSat)
			got, _ := inQtum.Float64()
			assert.Equal(test.Qtum, got)
		})
		t.Run(test.Name+": Convert From Qtum To Satoshi", func(t *testing.T) {
			inQtum := decimal.NewFromFloat(test.Qtum)
			inSat := ConvertFromQtumToSatoshis(inQtum)
			got, _ := inSat.Float64()
			assert.Equal(float64(test.Satoshi), got)
		})
		t.Run(test.Name+": Convert From Satoshi To Wei in Hex", func(t *testing.T) {
			if test.Satoshi == 0 {
				// cannot convert 0 satoshi to wei
				return
			}
			inSat := big.NewInt(int64(test.Satoshi))
			inWei := ConvertFromSatoshiToWei(inSat)
			got := EncodeBig(inWei)
			want := strings.ToLower(test.Wei)
			assert.Equal(want, got)
		})
		t.Run(test.Name+": Convert From Wei To Qtum", func(t *testing.T) {
			got, err := ConvertWeiToQtum(test.Wei)
			assert.NoError(err)
			want := float64(test.Qtum)
			assert.Equal(want, got)
		})

	}
}
