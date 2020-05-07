package apiret

type Return struct {
	RequestID string      `json:"request_id"`
	Result    interface{} `json:"result"`
}

func NewReturn(r interface{}, reqId string) *Return {
	return &Return{
		RequestID: reqId,
		Result:    r,
	}
}
