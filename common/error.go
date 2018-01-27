package common

type Error struct {
	ErrorType    string `json:"errorType"`
	IsError      bool   `json:"isError"`
	ErrorMessage string `json:"errorMessage"`
	TraceId      string `json:"traceId"`
}
