package ecode2

type Ecode2 interface {
	ToString() string
	ToInt() int
}

var (
	OK         = &Ecode{Code: 0, Message: "OK"}
	REQERROR   = &Ecode{Code: -400, Message: "参数错误"}
	INNERERROR = &Ecode{Code: -500, Message: "内部错误"}
)

// ok ecode
type Ecode struct {
	Code    int
	Message string
}

func (e *Ecode) ToInt() int {
	return e.Code
}

func (e *Ecode) ToString() string {
	return e.Message
}

func (e *Ecode) Error() string {
	return e.Message
}
