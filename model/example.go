package model

var (
	RetErr = &Ecode{
		Code:    10000,
		Message: "请求发生失败",
	}
)

var (
	StateOpen   = int64(1)
	StateClose  = int64(0)
	LevelLight  = int64(1) //轻
	LevelSevere = int64(2) //重
)
