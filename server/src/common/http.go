package common

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Tag  string      `json:"tag,omitempty"`
	Data interface{} `json:"data"`
}