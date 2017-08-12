package errcode

// result code

const (
	// common
	CODE_COMMON_SUCCESS     = 0 // 成功
	CODE_COMMON_ERROR       = 1 // 通用错误
	CODE_COMMON_EMPTY_ARG   = 2 // 参数为空
	CODE_COMMON_WRONG_ARG   = 3 // 错误的参数
	CODE_COMMON_ILLEGAL_ARG = 4 // 非法的参数
	CODE_COMMON_NOT_FOUND   = 5 // 数据未发现

	// usercenter
	CODE_UC_TOKEN_EXPIRED = 1000 // 登录失效

	// smartgate

	// gate
	CODE_GATE_INVALID_GATE = 3100 //闸机ID无效

	// verify evidence
	CODE_GATE_INVALID_EVIDENCE   = 3201 //无效凭证
	CODE_GATE_EXPIRED_EVIDENCE   = 3202 //凭证已过期
	CODE_GATE_NOT_MATCH_EVIDENCE = 3203 //凭证与机闸不匹配
	CODE_GATE_ROUTER_EXCEPTION   = 3204 //用户存在异常行程
	CODE_GATE_USER_NO_PAY        = 3205 //用户不符合付费标准

	//router
	CODE_GATE_MORE_IN  = 3301
	CODE_GATE_LATE_IN = 3302
	CODE_GATE_EARLY_OUT = 3303
	CODE_GATE_MORE_OUT = 3304
)

var msg_map = map[int]string{
	CODE_COMMON_SUCCESS:     "success",
	CODE_COMMON_ERROR:       "error",
	CODE_COMMON_EMPTY_ARG:   "required argument",
	CODE_COMMON_WRONG_ARG:   "wrong argument",
	CODE_COMMON_ILLEGAL_ARG: "illegal argument",
	CODE_COMMON_NOT_FOUND: "not found",

	CODE_UC_TOKEN_EXPIRED: "token was expired",

	CODE_GATE_INVALID_GATE:       "invalid gate id",
	CODE_GATE_INVALID_EVIDENCE:   "invalid evidence",
	CODE_GATE_EXPIRED_EVIDENCE:   "expired gate id",
	CODE_GATE_NOT_MATCH_EVIDENCE: "evidence not match the gate direction",
	CODE_GATE_USER_NO_PAY:        "user has no balance or quick pay",

	CODE_GATE_MORE_IN:  "[warning] created inbound router, but user has more inbound routers",
	CODE_GATE_LATE_IN: "[warning] user completed router later",
	CODE_GATE_EARLY_OUT: "[warning] created outbound router, but user has no inbound router",
	CODE_GATE_MORE_OUT: "[warning] created outbound router, but user has more early outbound routers",
}

func GetMsg(code int) string {
	return msg_map[code]
}
