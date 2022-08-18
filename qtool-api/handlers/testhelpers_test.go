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
