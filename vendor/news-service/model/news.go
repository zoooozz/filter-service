package model

import (
	"encoding/xml"
	"golang-kit/time"
)

// member info
type MemberInfo struct {
	Base    *MemberBaseInfo
	Service *MemberServiceInfo
	Company *MemberCompanyInfo
	Support *MemberSupportInfo
}

// member_base_info
type MemberBaseInfo struct {
	Id       int64     `json:"id"`
	Account  string    `json:"account"`
	Password string    `json:"password"`
	Sex      int8      `json:"sex"`
	Avatar   string    `json:"avatar"`
	Status   int8      `json:"status"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

// member_service_info
type MemberServiceInfo struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`
	Level     int8      `json:"level"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

// member_company_info
type MemberCompanyInfo struct {
	Id              int64     `json:"id"`
	Uid             int64     `json:"uid"`
	CompanyName     string    `json:"company_name"`
	BusinessSection string    `json:"business_section"`
	LicensePic      string    `json:"license_pic"`
	ContactName     string    `json:"contact_name"`
	Mobile          int64     `json:"mobile"`
	Phone           int64     `json:"phone"`
	QQ              int64     `json:"qq"`
	Email           string    `json:"email"`
	Addr            string    `json:"addr"`
	Website         string    `json:"website"`
	WechatPic       string    `json:"wechat_pic"`
	Scope           string    `json:"scope"`
	Brief           string    `json:"brief"`
	Banner          string    `json:"banner"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
}

// member_support_info
type MemberSupportInfo struct {
	Id      int64     `json:"id"`
	Uid     int64     `json:"uid"`
	Name    string    `json:"name"`
	Mobile  int64     `json:"mobile"`
	Phone   int64     `json:"phone"`
	Wechat  string    `json:"wechat"`
	QQ      int64     `json:"qq"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// template
type Template struct {
	Id        int64     `json:"id"`
	Uid       int64     `json:"uid"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Cover     string    `json:"cover"`
	Paragraph string    `json:"paragraph"`
	Sentence  string    `json:"sentence"`
	Count     int64     `json:"count"`
	Status    int8      `json:"status"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}

// business task
type BusinessClockTask struct {
	Task     *ClockTask `json:"task"`
	Business *Business  `json:"business"`
}

// business task
type BusinessSingleTask struct {
	Task     *SingleTask
	Business *Business
}

// business
type Business struct {
	Id           int64     `json:"id"`
	Uid          int64     `json:"uid"`
	IsBatch      int8      `json:"is_batch"`
	PubType      int8      `json:"pub_type"`
	Section      string    `json:"section"`   // 栏目
	Brand        string    `json:"brand"`     // 品牌
	Title        string    `json:"title"`     // 所标题
	TitleNum     int64     `json:"title_num"` // 标题数量
	Cover        string    `json:"cover"`
	Price        int64     `json:"price"`
	Num          int64     `json:"num"`
	Addr         string    `json:"addr"`
	SupplyCount  int64     `json:"supply_count"`
	SendDeadline int64     `json:"send_deadline"`
	TemplateIds  string    `json:"template_ids"`
	PubTime      string    `json:"pub_time"`
	DayMaxNum    int64     `json:"day_max_num"`
	Content      string    `json:"content"`
	KeyWord      string    `json:"keyword"`
	Status       int8      `json:"status"`
	Created      time.Time `json:"created"`
	Updated      time.Time `json:"updated"`
	CreatedDate  string    `json:"created_date"`
	UpdatedDate  string    `json:"updated_date"`
}

// task
type ClockTask struct {
	Id         int64     `json:"id"`
	Uid        int64     `json:"uid"`
	BusinessId int64     `json:"business_id"`
	Count      int64     `json:"count"`
	Status     int8      `json:"status"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}

// single task
type SingleTask struct {
	Id          int64     `json:"id"`
	Uid         int64     `json:"uid"`
	ClockTaskId int64     `json:"clock_task_id"`
	BusinessId  int64     `json:"business_id"`
	Title       string    `json:"title"`
	Status      int8      `json:"status"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

// business_sended
type BusinessSended struct {
	Id          int64     `json:"id"`
	Uid         int64     `json:"uid"`
	Section     string    `json:"section"` // 栏目
	BusinessId  int64     `json:"business_id"`
	ClockTaskId int64     `json:"clock_task_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Status      int8      `json:"status"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type SmsInfo struct {
	Returnsms     xml.Name `xml:"returnsms"`
	Returnstatus  string   `xml:"returnstatus"`
	Message       string   `xml:"message"`
	Remainpoint   int64    `xml:"remainpoint"`
	TaskID        string   `xml:"taskID"`
	SuccessCounts int64    `xml:"successCounts"`
}
