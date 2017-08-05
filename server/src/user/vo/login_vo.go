package vo

type WxappLogin struct {
	Code string `json:"code"`
}

type WxappLoginToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type WxappSession struct {
	Session_key string `json:"session_key"`
	ExpiresIn   int    `json:"expires_in"`
	Openid      string `json:"openid"`
	Unionid     string `json:"unionid"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
}

type VerifyToken struct {
	Token string `json:"token"`
}
