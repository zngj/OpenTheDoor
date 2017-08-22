package errcode

type SGError struct {
	code int
	msg  string
}

func NewError(code int) *SGError {
	return NewError2(code, GetMsg(code))
}

func NewError2(code int, msg string) *SGError {
	sge := new(SGError)
	sge.code = code
	sge.msg = msg
	return sge
}

func (sge *SGError) Code() int {
	return sge.code
}

func (sge *SGError) Error() string {
	return sge.msg
}

var (
	SGErrWrongArg     = NewError(CODE_COMMON_WRONG_ARG)
	SGErrMobileDuplicate       = NewError(CODE_UC_MOBILE_DUPLICATE)
	SGErrMobileNotFound        = NewError(CODE_UC_MOBILE_NOT_FOUND)
	SGErrWrongPassword         = NewError(CODE_UC_WRONG_PASSWORD)
	SGErrInvalidWxCode         = NewError(CODE_UC_INVALID_WXCODE)
	SGErrNotMatchGateDirection = NewError(CODE_GATE_NOT_MATCH_EVIDENCE)
	SGErrMoreIn                = NewError(CODE_GATE_MORE_IN)
	SGErrExistIn               = NewError(CODE_GATE_EXIST_IN)
	SGErrDiffIn                = NewError(CODE_GATE_DIFF_IN)
	SGErrLateIn                = NewError(CODE_GATE_LATE_IN)
	SGErrEarlyOut              = NewError(CODE_GATE_EARLY_OUT)
	SGErrMoreOut               = NewError(CODE_GATE_MORE_OUT)
)
