package vo

type WalletVO struct {
	Balance    float32 `json:"balance"`
	WxpayQuick bool    `json:"wxpay_quick"`
}

type WalletChargeVO struct {
	Money float32 `json:"money"`
}
