package model

import (
	"golang-kit/ecode2"
)

var (
	// 未发送验证码
	NotSendCaptcha = &ecode2.Ecode{
		Code:    1,
		Message: "未发送验证码",
	}
	// 验证码错误
	CaptchaIsError = &ecode2.Ecode{
		Code:    2,
		Message: "验证码错误",
	}
	// 发送验证码失败
	CaptchaSendFail = &ecode2.Ecode{
		Code:    3,
		Message: "发送验证码失败",
	}
)
