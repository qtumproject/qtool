package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/qtumproject/qtool/qtool-api/types"
	"github.com/stretchr/testify/assert"
)

const (
	addressHex    = "7926223070547d2d15b2ef5e7383e541c338ffe9"
	addressBase58 = "qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW"
	privateKeyWIF = "cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk"
	privateKeyHex = "00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35"
	pubKey        = "0299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112" // compressed

)

func assertHandlerResponse(t *testing.T, cmd func(c echo.Context) error, body *types.JSONRPCRequest, want *types.JSONRPCResponse) {
	t.Helper()
	bodyBytes, _ := json.Marshal(body)
	bodyBuf := bytes.NewBuffer(bodyBytes)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bodyBuf)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, cmd(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var got *types.JSONRPCResponse
		err := json.Unmarshal(rec.Body.Bytes(), &got)
		if err != nil {
			t.Errorf("Error unmarshalling JSONRPCResponse: %s", err)
		}
		assert.Equal(t, want, got)
	}

}

func getJsonRpcRequestBody(method, data, format, blockchain, network string) *types.JSONRPCRequest {
	params := types.RequestParams{
		Format:     format,
		Data:       data,
		Blockchain: blockchain,
		Network:    network,
	}
	paramsBytes, _ := json.Marshal(params)

	body := types.JSONRPCRequest{
		ID:     1,
		Method: method,
		Params: paramsBytes,
	}

	return &body
}
