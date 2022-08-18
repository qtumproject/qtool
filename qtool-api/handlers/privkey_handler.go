package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alejoacosta74/qtool/pkg/tools"
	"github.com/alejoacosta74/qtool/qtool-api/types"

	"github.com/labstack/echo/v4"
)

func PrivateKeyHandler(c echo.Context) error {
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
	case "convertprivkey":
		if params.Format == "b58" {
			result, err = tools.ConvertPrivateKeyToHex(params.Data)
		}
		if params.Format == "hex" {
			result, err = tools.ConvertPrivateKeyToWIF(params.Data, params.Network)
		}
	case "getaddrfromprivkey":
		result, err = tools.GetAddressFromPrivkey(params.Data, params.Blockchain, params.Network, params.Format)
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
