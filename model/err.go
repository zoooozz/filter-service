package model

import (
	"net/http"
)

type Ecode struct {
	Code    int
	Message string
}

var (
	Notdefinition = &Ecode{
		Code:    404,
		Message: "未定义",
	}

	Success = &Ecode{
		Code:    0,
		Message: "请求成功",
	}

	Notdefinitionparams = &Ecode{
		Code:    400,
		Message: "参数错误",
	}
)

func OutputRet(e *Ecode, result ...map[string]interface{}) (code int, message map[string]interface{}) {

	code = http.StatusOK
	message = make(map[string]interface{})

	if result != nil {
		message["data"] = result[0]
	}

	message["code"] = e.Code
	message["msg"] = e.Message
	return
}
