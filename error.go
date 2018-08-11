package sbankensdk

type sbankenError struct {
	ErrorType    *int    `json:"errorType,omitempty"`
	IsError      *bool   `json:"isError,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	TraceID      *string `json:"traceId,omitempty"`
}
