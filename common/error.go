package common

type Error struct {
	ErrorType    int    `json:"errorType"`
	IsError      bool   `json:"isError"`
	ErrorMessage string `json:"errorMessage"`
	TraceId      string `json:"traceId"`
}
