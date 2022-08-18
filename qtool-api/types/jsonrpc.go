package types

import "encoding/json"

const (
	jsonrpcVersion = "2.0"
)

// http://www.jsonrpc.org/specification#request_object
type JSONRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      int             `json:"id"`
}

// http://www.jsonrpc.org/specification#response_object
type JSONRPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	// Error   string `json:"error"`
	Error *JSONRPCError `json:"error"`
	ID    int           `json:"id"`
}

// http://www.jsonrpc.org/specification#error_object
type JSONRPCError struct {
	JSONRPC string `json:"jsonrpc"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	ID      int    `json:"id"`
}

func NewJSONRPCResponse(id int, result json.RawMessage, err *JSONRPCError) *JSONRPCResponse {
	return &JSONRPCResponse{
		JSONRPC: jsonrpcVersion,
		Result:  result,
		Error:   err,
		ID:      id,
	}
}

func NewJSONRPCResponseError(id, code int, message, data string) *JSONRPCError {
	return &JSONRPCError{
		JSONRPC: jsonrpcVersion,
		ID:      id,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type RequestParams struct {
	Data       string `json:"data"`
	Format     string `json:"format"`
	Network    string `json:"network"`
	Blockchain string `json:"blockchain"`
}
