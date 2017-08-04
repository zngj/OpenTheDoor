package httpx

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	Input  string      `json:"input,omitempty"`
}
