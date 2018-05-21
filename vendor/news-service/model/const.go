package model

const (
	// 发布类型
	SUPPLYTYPE = int8(0) // 供求类型
	NEWSTYPE   = int8(1) // 新闻类型

	// 是否是批量
	ISBATCH    = int8(0) // 是
	ISNOTBATCH = int8(1) // 不是

	// 当日发送最大数量
	DAYMAXPUBNUMBER = int64(5000)

	// Clock_Task 定时任务状态定义
	ClockTask_complete   = int8(20) // 完成
	ClockTask_del        = int8(3)  // 删除
	ClockTask_stop       = int8(2)  // 暂停
	ClockTask_doing      = int8(1)  // 正在运行
	ClockTask_uncomplete = int8(0)  // 未完成(还未开始)

	// Single_Task 单个任务状态定义
	SingleTask_complete   = int8(20) // 已完成
	SingleTask_uncomplete = int8(0)  // 未完成(还未开始)

	// template status
	Template_normal = int8(0) // 正常
	Template_del    = int8(1) // 删除

	// business status
	BusinessUnVerify  = int8(0) // 未审核
	BusinessHasVerify = int8(1) // 已审核
	BusinessDel       = int8(2) // 删除
	BusinessRefuse    = int8(3) // 拒审核
)
