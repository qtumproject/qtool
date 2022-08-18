package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alejoacosta74/qtool/pkg/tools"
	"github.com/alejoacosta74/qtool/qtool-api/types"
	"github.com/labstack/echo/v4"
)

func ScriptPubKeyHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	var req types.JSONRPCRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var result interface{}
	var params types.RequestParams
	err = json.Unmarshal(req.Params, &params)
	if err != nil {
		return responseJsonErr(c, req.ID, err)
	}
	switch req.Method {
	case "getaddressfromscriptpubKey":
		result, err = tools.AddressFromP2PKScript(params.Data, params.Blockchain, params.Network)
	default:
		return responseJsonErr(c, req.ID, fmt.Errorf("method not found"))
	}
	if err != nil {
		return responseJsonErr(c, req.ID, err)
	}
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return responseJsonErr(c, req.ID, err)
	}
	return responseJsonResult(c, req.ID, resultBytes)
}
