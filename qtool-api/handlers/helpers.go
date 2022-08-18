package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qtumproject/qtool/qtool-api/types"
)

// creates and sends a new JSONRPCResponse object with the given error
func responseJsonErr(c echo.Context, reqID int, err error) error {
	jsonError := types.NewJSONRPCResponseError(reqID, -1, err.Error(), "")
	resp := types.NewJSONRPCResponse(reqID, nil, jsonError)
	fmt.Printf("Returning error: %+v\n", resp)
	return c.JSON(http.StatusBadRequest, resp)
}

// creates and sends a new JSONRPCResponse object with the given result
func responseJsonResult(c echo.Context, reqID int, result []byte) error {
	resp := types.NewJSONRPCResponse(reqID, result, nil)
	// fmt.Printf("Returning resp: %+v\n", resp)
	return c.JSON(http.StatusOK, resp)
}
