package handlers

import (
	"encoding/json"
	"testing"

	"github.com/qtumproject/qtool/pkg/tools"
	"github.com/qtumproject/qtool/qtool-api/types"
)

func TestHandlerPrivKeyConvert(t *testing.T) {

	t.Run("ConvertPrivKey to Hex", func(t *testing.T) {
		body := getJsonRpcRequestBody("convertprivkey", privateKeyWIF, "b58", "qtum", "testnet")
		result := &tools.ConvertPrivateKeyResult{
			PrivateKey: privateKeyHex,
		}
		resultBytes, _ := json.Marshal(result)
		want := types.NewJSONRPCResponse(1, resultBytes, nil)
		assertHandlerResponse(t, PrivateKeyHandler, body, want)
	})
	t.Run("ConvertPrivKey to WIF", func(t *testing.T) {
		body := getJsonRpcRequestBody("convertprivkey", privateKeyHex, "hex", "qtum", "testnet")
		result := &tools.ConvertPrivateKeyResult{
			PrivateKey: privateKeyWIF,
		}
		resultBytes, _ := json.Marshal(result)
		want := types.NewJSONRPCResponse(1, resultBytes, nil)
		assertHandlerResponse(t, PrivateKeyHandler, body, want)
	})
}

func TestHandlerGetAddrFromPrivKey(t *testing.T) {
	t.Run("Get address from WIF priv key", func(t *testing.T) {
		body := getJsonRpcRequestBody("getaddrfromprivkey", privateKeyWIF, "b58", "qtum", "testnet")
		result := &tools.GetAddressFromPrivkeyResult{
			PrivateKeyHex: privateKeyHex,
			PrivateKeyWIF: privateKeyWIF,
			PublicKeyHex:  pubKey,
			AddressHex:    addressHex,
			AddressBase58: addressBase58,
		}

		resultBytes, _ := json.Marshal(result)
		want := types.NewJSONRPCResponse(1, resultBytes, nil)

		assertHandlerResponse(t, PrivateKeyHandler, body, want)
	})
}

// func assertConvertPrivKey(t *testing.T, body types.ConvertKeyJSONRPCRequest, want types.JSONRPCResponse) {
// 	t.Helper()
// 	bodyBytes, _ := json.Marshal(body)
// 	bodyBuf := bytes.NewBuffer(bodyBytes)

// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", bodyBuf)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	// Assertions
// 	if assert.NoError(t, PrivateKey(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)

// 		got := types.JSONRPCResponse{}
// 		err := json.Unmarshal(rec.Body.Bytes(), &got)
// 		if err != nil {
// 			t.Errorf("Error unmarshalling JSONRPCResponse: %s", err)
// 		}
// 		assert.Equal(t, want, got)
// 	}

// }

// func assertConvertPrivKey(t *testing.T, body types.ConvertKeyJSONRPCRequest, want types.JSONRPCResponse) {
// 	t.Helper()
// 	bodyBytes, _ := json.Marshal(body)
// 	bodyBuf := bytes.NewBuffer(bodyBytes)

// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", bodyBuf)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	// Assertions
// 	if assert.NoError(t, PrivateKeyHandler(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)

// 		got := types.JSONRPCResponse{}
// 		err := json.Unmarshal(rec.Body.Bytes(), &got)
// 		if err != nil {
// 			t.Errorf("Error unmarshalling JSONRPCResponse: %s", err)
// 		}
// 		assert.Equal(t, want, got)
// 	}

// }
