package handlers

import (
	"encoding/json"
	"testing"

	"github.com/qtumproject/qtool/pkg/tools"
	"github.com/qtumproject/qtool/qtool-api/types"
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
