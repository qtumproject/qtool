package common

import (
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

// Converts a satoshis to qtum balance
func ConvertFromSatoshisToQtum(inSatoshis decimal.Decimal) decimal.Decimal {
	return inSatoshis.Div(decimal.NewFromFloat(float64(1e8)))
}

// Converts a qtum balance to satoshis
func ConvertFromQtumToSatoshis(inQtum decimal.Decimal) decimal.Decimal {
	return inQtum.Mul(decimal.NewFromFloat(float64(1e8)))
}

// Converts a satoshis to wei balance
func ConvertFromSatoshiToWei(inSatoshis *big.Int) *big.Int {
	return inSatoshis.Mul(inSatoshis, big.NewInt(1e10))
}

// Converts a hex string representing a value in wei to a value in satoshis
func ConvertHexFromSatoshiToWei(inSatoshis string) (string, error) {
	hasPrefix := true
	if !strings.HasPrefix(inSatoshis, "0x") {
		inSatoshis = "0x" + inSatoshis
		hasPrefix = false
	}
	valSatoshis, err := hexutil.DecodeBig(inSatoshis)
	if err != nil {
		return "", errors.Wrap(err, "error decoding satoshis")
	}

	valWei := ConvertFromSatoshiToWei(valSatoshis)
	inWei := hexutil.EncodeBig(valWei)
	if !hasPrefix {
		inWei = strings.TrimPrefix(inWei, "0x")
	}
	return inWei, nil
}
