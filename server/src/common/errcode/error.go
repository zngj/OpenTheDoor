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
	SGErrMoreIn = NewError(CODE_GATE_MORE_IN)
	SGErrLateIn = NewError(CODE_GATE_LATE_IN)
	SGErrEarlyOut = NewError(CODE_GATE_EARLY_OUT)
	SGErrMoreOut = NewError(CODE_GATE_MORE_OUT)
)
