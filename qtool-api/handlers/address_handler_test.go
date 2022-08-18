package handlers

import (
	"encoding/json"
	"testing"

	"github.com/alejoacosta74/qtool/pkg/tools"
	"github.com/alejoacosta74/qtool/qtool-api/types"
)

const (
	addressHex    = "7926223070547d2d15b2ef5e7383e541c338ffe9"
	addressBase58 = "qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW"
)

func TestHandlerAddressConvert(t *testing.T) {

	t.Run("Convert Address to Hex", func(t *testing.T) {

		body := getJsonRpcRequestBody("convertaddress", addressBase58, "b58", "qtum", "testnet")
		result := &tools.ConvertAddressResult{
			Address: addressHex,
		}
		resultBytes, _ := json.Marshal(result)
		want := types.NewJSONRPCResponse(1, resultBytes, nil)
		assertHandlerResponse(t, AddressHandler, body, want)
	})

	t.Run("Convert Address to Base58", func(t *testing.T) {

		body := getJsonRpcRequestBody("convertaddress", addressHex, "hex", "qtum", "testnet")
		result := &tools.ConvertAddressResult{
			Address: addressBase58,
		}
		resultBytes, _ := json.Marshal(result)
		want := types.NewJSONRPCResponse(1, resultBytes, nil)
		assertHandlerResponse(t, AddressHandler, body, want)
	})
}
