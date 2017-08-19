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
	CODE_UC_TOKEN_EXPIRED    = 1000 // 登录失效
	CODE_UC_TOKEN_REQUIRED   = 1000 // 登录失效
	CODE_UC_MOBILE_DUPLICATE = 1001 //注册名重复
	CODE_UC_MOBILE_NOT_FOUND = 1002 //手机号码不存在
	CODE_UC_WRONG_PASSWORD   = 1003 //密码不正确

	CODE_UC_INVALID_WXCODE   = 40029

	// smartgate

	// gate
	CODE_GATE_INVALID_GATE = 3100 //闸机ID无效

	// verify evidence
	CODE_GATE_INVALID_EVIDENCE   = 3201 //无效凭证
	CODE_GATE_USED_EVIDENCE      = 3202 //凭证已使用
	CODE_GATE_EXPIRED_EVIDENCE   = 3203 //凭证已过期
	CODE_GATE_NOT_MATCH_EVIDENCE = 3204 //凭证与机闸不匹配
	CODE_GATE_ROUTER_NO_IN       = 3205 //用户存在异常行程
	CODE_GATE_ROUTER_EXCEPTION   = 3206 //用户存在异常行程
	CODE_GATE_USER_NO_PAY        = 3207 //用户不符合付费标准

	//router
	CODE_GATE_MORE_IN   = 3301
	CODE_GATE_EXIST_IN  = 3302
	CODE_GATE_DIFF_IN   = 3303
	CODE_GATE_LATE_IN   = 3304
	CODE_GATE_EARLY_OUT = 3305
	CODE_GATE_MORE_OUT  = 3306
	CODE_GATE_DIFF_OUT  = 3307
)

var msg_map = map[int]string{
	//CODE_COMMON_SUCCESS:     "success",
	//CODE_COMMON_ERROR:       "error",
	//CODE_COMMON_EMPTY_ARG:   "required argument",
	//CODE_COMMON_WRONG_ARG:   "wrong argument",
	//CODE_COMMON_ILLEGAL_ARG: "illegal argument",
	//CODE_COMMON_NOT_FOUND:   "not found",
	//
	//CODE_UC_TOKEN_EXPIRED: "token was expired",
	//
	//CODE_GATE_INVALID_GATE:       "invalid gate id",
	//CODE_GATE_INVALID_EVIDENCE:   "invalid evidence",
	//CODE_GATE_USED_EVIDENCE:      "used evidence",
	//CODE_GATE_EXPIRED_EVIDENCE:   "expired gate id",
	//CODE_GATE_NOT_MATCH_EVIDENCE: "evidence not match the gate direction",
	//CODE_GATE_ROUTER_NO_IN:       "user has no inbound service",
	//CODE_GATE_USER_NO_PAY:        "user has no balance or quick pay",
	//
	//CODE_GATE_MORE_IN:   "[warning] created inbound router, but user has more inbound routers",
	//CODE_GATE_EXIST_IN:  "existed inbound router for a long time, please get out first",
	//CODE_GATE_DIFF_IN:   "can't get in from different station",
	//CODE_GATE_LATE_IN:   "[warning] user completed router later",
	//CODE_GATE_EARLY_OUT: "[warning] created outbound router, but user has no inbound router",
	//CODE_GATE_MORE_OUT:  "[warning] created outbound router, but user has more early outbound routers",

	CODE_COMMON_SUCCESS:     "成功",
	CODE_COMMON_ERROR:       "错误",
	CODE_COMMON_EMPTY_ARG:   "参数不能为空",
	CODE_COMMON_WRONG_ARG:   "错误参数",
	CODE_COMMON_ILLEGAL_ARG: "非法参数",
	CODE_COMMON_NOT_FOUND:   "没有数据",

	CODE_UC_TOKEN_EXPIRED:     "登录失效",
	CODE_UC_MOBILE_DUPLICATE:  "手机号已被注册",
	CODE_UC_MOBILE_NOT_FOUND : "手机号码不存在",
	CODE_UC_WRONG_PASSWORD:    "密码不正确",
	CODE_UC_INVALID_WXCODE:    "无效的微信授权临时票据(code)",

	CODE_GATE_INVALID_GATE:       "闸机编号无效",
	CODE_GATE_INVALID_EVIDENCE:   "无效的凭证",
	CODE_GATE_USED_EVIDENCE:      "凭证已使用",
	CODE_GATE_EXPIRED_EVIDENCE:   "凭证已过期",
	CODE_GATE_NOT_MATCH_EVIDENCE: "凭证与闸机不匹配",
	CODE_GATE_ROUTER_NO_IN:       "无进站行程",
	CODE_GATE_USER_NO_PAY:        "账户余额不足",

	CODE_GATE_MORE_IN:   "多人进站",
	CODE_GATE_EXIST_IN:  "无效进站，请先出站",
	CODE_GATE_DIFF_IN:   "不允许在多个站点同时进站",
	CODE_GATE_LATE_IN:   "进站数据晚于出站数据",
	CODE_GATE_EARLY_OUT: "出站数据早于进站数据",
	CODE_GATE_MORE_OUT:  "多人出站",
}

func GetMsg(code int) string {
	return msg_map[code]
}
