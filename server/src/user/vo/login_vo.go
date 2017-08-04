package vo

type WxappLogin struct {
	Code string `json:"code"`
}

type WxappLoginToken struct {
	Token string `json:"token"`
}

type WxappSession struct {
	Openid      string `json:"openid"`
	Session_key string `json:"session_key"`
	Unionid     string `json:"unionid"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

type VerifyToken struct {
	Token string `json:"token"`
}
