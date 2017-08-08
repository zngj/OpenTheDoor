package errcode

// result code

const (
	// common
	CODE_COMMON_SUCCESS   = 0  // 成功
	CODE_COMMON_ERROR     = 1 // 通用错误
	CODE_COMMON_EMPTY_ARG = 2 // 参数为空
	CODE_COMMON_WRONG_ARG = 3 // 错误的参数
	CODE_COMMON_ILLEGAL_ARG = 4 // 非法的参数

	// usercenter
	CODE_UC_TOKEN_EXPIRED = 1000 // 登录失效

	// smartgate

	// gate
	CODE_GATE_NO_GATE = 3100 //闸机ID不存在
	CODE_GATE_NO_EVIDENCE = 3201 //凭证不存在
	CODE_GATE_EXPIRED_EVIDENCE = 3202 //凭证已过期
	CODE_GATE_NOT_MATCH_EVIDENCE = 3203 //凭证与机闸不匹配
	CODE_GATE_USER_NO_PAY = 3204 //用户不符合付费标准

)

var msg_map = map[int]string {
	CODE_COMMON_SUCCESS :"success",
	CODE_COMMON_ERROR :"error",
	CODE_COMMON_EMPTY_ARG :"required argument",
	CODE_COMMON_WRONG_ARG :"wrong argument",
	CODE_COMMON_ILLEGAL_ARG :"illegal argument",

	CODE_UC_TOKEN_EXPIRED: "token was expired",
}

func GetMsg(code int) string {
	return msg_map[code]
}