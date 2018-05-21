package ecode

import (
	"strconv"
)

type ecode int

func (e ecode) Error() string {
	return strconv.FormatInt(int64(e), 10)
}

func (e ecode) Message() (str string) {
	var ok bool
	if str, ok = ecodeMap[e]; ok {
		return
	} else {
		str = "未知错误"
	}
	return
}

func To(i int) error {
	return ecode(i)
}

func From(e error) ecode {
	if e == nil {
		return OK
	}
	i, err := strconv.ParseInt(e.Error(), 10, 64)
	if err != nil {
		return ServerErr
	}
	return ecode(i)
}

// error to ecode
func Lookup(e error) ecode {
	return From(e)
}
