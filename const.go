package xmrto

import (
	"fmt"
)

// APIBaseAddress is the base URL of the xmr.to API
const APIBaseAddress = "https://xmr.to/api"

// APITestnetBaseAddress is the base URL of the xmr.to TESTNET API
const APITestnetBaseAddress = "https://test.xmr.to/api/"

// APIVersion is the API version identifier (currently v3)
const APIVersion = "v3"

// APIConversionDirection is the API conversion direction (currently only xmr2btc)
const APIConversionDirection = "xmr2btc"

// ErrorCode is a xmr.to error code table
type ErrorCode int

// APIError represents an error message by the xmr.to API
type APIError struct {
	APIError        string `json:"error"`
	APIErrorMessage string `json:"error_msg"`
}

func (we *APIError) Error() string {
	return fmt.Sprintf("%v: (%v) - See https://xmrto-api.readthedocs.io/en/latest/ for help!", we.APIError, we.APIErrorMessage)
}
