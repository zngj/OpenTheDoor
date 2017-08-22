package vo

type SignUpVo struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	VerifyCode  string `json:"verify_code"`
}

type LoginVo struct {
	PhoneNumber   string `json:"phone_number"`
	Password string `json:"password"`
}

type WxappLogin struct {
	Code string `json:"code"`
}

type LoginToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type WxSession struct {
	Client       string `json:"client"`
	AccessToken  string `json:"access_token"`
	SessionKey   string `json:"session_key"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Unionid      string `json:"unionid"`
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

type VerifyToken struct {
	AccessToken string `json:"access_token"`
}
