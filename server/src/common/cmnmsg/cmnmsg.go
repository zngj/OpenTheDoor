package cmnmsg

// common code between 0 and -999

const (
	CODE_COMMON_SUCCESS   = 0  // 成功
	CODE_COMMON_ERROR     = 1 // 通用错误
	CODE_COMMON_EMPTY_ARG = 2 // 参数为空
	CODE_COMMON_WRONG_ARG = 3 // 错误的参数
	CODE_COMMON_ILLEGAL_ARG = 4 // 非法的参数
)

var msg_map = map[int]string {
	CODE_COMMON_SUCCESS :"success",
	CODE_COMMON_ERROR :"error",
	CODE_COMMON_EMPTY_ARG :"required argument",
	CODE_COMMON_WRONG_ARG :"wrong argument",
	CODE_COMMON_ILLEGAL_ARG :"illegal argument",
}

func GetMsg(code int) string {
	return msg_map[code]
}

func AddMsg(msgs map[int]string) {
	for k, v := range msgs {
		msg_map[k] = v
	}
}
