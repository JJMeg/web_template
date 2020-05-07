package errors

import "fmt"

type ErrorResponse struct {
	RequestID string `json:"request_id"`
	Err       Error  `json:"err"`
}

func NewErrorResponse(reqId string, err Error) *ErrorResponse {
	return &ErrorResponse{
		RequestID: reqId,
		Err:       err,
	}
}

func (r *ErrorResponse) GetStatusCode() int {
	return r.Err.Code
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("[ %s ]: { %d , %s }", r.RequestID, r.Err.Code, r.Err.Message)
}
