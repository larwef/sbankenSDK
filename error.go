package sbankenSDK

type sbankenError struct {
	ErrorType    int    `json:"errorType,omitempty"`
	IsError      bool   `json:"isError,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	TraceId      string `json:"traceId,omitempty"`
}
