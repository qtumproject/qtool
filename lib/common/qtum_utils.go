package common

import (
	"math/big"
	"strings"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

// Converts a satoshis balance to qtum balance
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

// Converts a hex string representing a value in wei to a value in satoshis
func ConvertHexFromWeiToSatoshi(inWei string) (string, error) {
	hasPrefix := true
	if !strings.HasPrefix(inWei, "0x") {
		inWei = "0x" + inWei
		hasPrefix = false
	}
	valWei, err := hexutil.DecodeBig(inWei)
	if err != nil {
		return "", errors.Wrap(err, "error decoding wei")
	}

	valSatoshis := valWei.Div(valWei, big.NewInt(1e10))
	inSatoshis := hexutil.EncodeBig(valSatoshis)
	if !hasPrefix {
		inSatoshis = strings.TrimPrefix(inSatoshis, "0x")
	}
	return inSatoshis, nil
}

// Converts a hex string representing a value in wei to a
// float representing the value in qtum
func ConvertWeiToQtum(inWei string) (float64, error) {
	inSat, err := ConvertHexFromWeiToSatoshi(inWei)
	if err != nil {
		return 0, err
	}
	inSatoshis, err := hexutil.DecodeBig(inSat)
	if err != nil {
		return 0, err
	}
	inQtum := ConvertFromSatoshisToQtum(decimal.NewFromBigInt(inSatoshis, 0))
	result, _ := inQtum.Float64()
	return result, nil
}
