package ecode

// define your code here
const (
	// common error code
	OK              ecode = 0
	AppKeyInvalid   ecode = -1   // 应用程序不存在或已被封禁
	AccessKeyErr    ecode = -2   // Access Key错误
	SignCheckErr    ecode = -3   // API校验密匙错误
	NoLogin         ecode = -101 // 账号未登录
	UserDisabled    ecode = -102 // 账号被封停
	CaptchaErr      ecode = -105 // 验证码错误
	UserInactive    ecode = -106 // 账号未激活
	UserNoMember    ecode = -107 // 账号非正式会员或在适应期
	AppDenied       ecode = -108 // 应用不存在或者被封禁
	MobileNoVerfiy  ecode = -110 // 未绑定手机
	CsrfNotMatchErr ecode = -111 // csrf 校验失败
	ServiceUpdate   ecode = -112 // 系统升级中
	LoginInfoErr    ecode = -114 // cookie信息有误
	UserNotExist    ecode = -115 // 用户不存在
	LoginFailed     ecode = -116 // 登录失败
	LogOutFailed    ecode = -117 // 退出失败
	Logined         ecode = -118 // 已登录
	UserExist       ecode = -119 // 用户已存在

	RequestErr ecode = -400 // 参数错误
	ServerErr  ecode = -500 // 服务器错误

	// reply
	ReplyNotExist ecode = 10000
)

var (
	ecodeMap = map[ecode]string{
		// base ecode
		OK:              "ok",
		ServerErr:       "服务器错误",
		RequestErr:      "参数错误",
		AppKeyInvalid:   "应用程序不存在或已被封禁",
		AccessKeyErr:    "Access Key错误",
		SignCheckErr:    "API校验密匙错误",
		NoLogin:         "账号未登录",
		UserDisabled:    "账号被封停",
		UserInactive:    "账号未激活",
		UserNoMember:    "账号非正式会员或在适应期",
		AppDenied:       "应用不存在或者被封禁",
		MobileNoVerfiy:  "未绑定手机",
		CsrfNotMatchErr: "csrf 校验失败",
		ServiceUpdate:   "系统升级中",
		LoginInfoErr:    "cookie信息有误",
		UserNotExist:    "用户不存在或者账号密码错误",
		LoginFailed:     "登录失败",
		LogOutFailed:    "退出失败",
		Logined:         "已登录",
		UserExist:       "用户已存在",

		// 评论
		ReplyNotExist: "评论不存在",
	}
)
