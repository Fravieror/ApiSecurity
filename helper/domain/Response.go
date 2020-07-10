package domain

type Response struct {
	Message string
	Code    string
	Data    []byte
	Errors  string
}

func NewResponse() *Response {
	return &Response{
		Message: "Ok",
		Code:    "200",
	}
}
