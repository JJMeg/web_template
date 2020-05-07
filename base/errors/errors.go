package errors

type Error struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (e Error) IsEmptyError() bool {
	return e.Code == 0 && e.Message == ""
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) String() string {
	return e.Status + ": " + e.Message
}

func (e Error) SerMsg(msg string) {
	e.Message = msg
}
