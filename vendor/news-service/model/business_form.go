package model

import (
	"github.com/mholt/binding"
	"golang-kit/ecode"
	"golang-kit/log"
	"net/http"
	"strings"
)

// add form
func (f *Business) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.IsBatch:      "is_batch",
		&f.PubType:      "pub_type",
		&f.Section:      "section",
		&f.Title:        "title",
		&f.Cover:        "cover",
		&f.Brand:        "brand",
		&f.Price:        "price",
		&f.Num:          "num",
		&f.Addr:         "addr",
		&f.SupplyCount:  "supply_count",
		&f.SendDeadline: "send_deadline",
		&f.TemplateIds:  "template_ids",
		&f.PubTime:      "pub_time",
		&f.DayMaxNum:    "day_max_num",
		&f.Content:      "content",
		&f.KeyWord:      "keyword",
	}
}

func (f *Business) Validate(req *http.Request) (err error) {
	if f.IsBatch != int8(ISBATCH) && f.IsBatch != int8(ISNOTBATCH) {
		err = ecode.RequestErr
		log.Error("IsBatch(%v) is not validate", f.IsBatch)
		return
	}
	if f.PubType != int8(SUPPLYTYPE) && f.PubType != int8(NEWSTYPE) {
		err = ecode.RequestErr
		log.Error("PubType(%v) is not validate", f.PubType)
		return
	}
	if f.Section == "" {
		err = ecode.RequestErr
		log.Error("Section(%v) is not validate", f.Section)
		return
	}
	if f.Title == "" {
		err = ecode.RequestErr
		log.Error("Title(%v) is not validate", f.Title)
		return
	}
	f.TitleNum = int64(strings.Count(f.Title, "\r\n") + 1)
	if f.IsBatch == ISNOTBATCH && f.PubType == SUPPLYTYPE {
		if f.Cover == "" {
			err = ecode.RequestErr
			log.Error("Cover(%v) is not validate", f.Cover)
			return
		}
	}
	// 供求信息的单独认证
	if f.PubType == SUPPLYTYPE {
		if f.Brand == "" {
			err = ecode.RequestErr
			log.Error("Brand(%v) is not validate", f.Brand)
			return
		}
		if f.Price == 0 {
			err = ecode.RequestErr
			log.Error("Price(%v) is not validate", f.Price)
			return
		}
		if f.Num == 0 {
			err = ecode.RequestErr
			log.Error("Num(%v) is not validate", f.Num)
			return
		}
		if f.Addr == "" {
			err = ecode.RequestErr
			log.Error("Addr(%v) is not validate", f.Addr)
			return
		}
		if f.SupplyCount == 0 {
			err = ecode.RequestErr
			log.Error("SupplyCount(%v) is not validate", f.SupplyCount)
			return
		}
		if f.SendDeadline == 0 {
			err = ecode.RequestErr
			log.Error("SendDeadline(%v) is not validate", f.SendDeadline)
			return
		}
	}
	// 批量发布要验证内容模板
	if f.IsBatch == ISBATCH {
		if f.TemplateIds == "" {
			err = ecode.RequestErr
			log.Error("TemplateIds(%v) is not validate", f.TemplateIds)
			return
		}
	}
	if f.PubTime == "" {
		err = ecode.RequestErr
		log.Error("PubTime(%v) is not validate", f.PubTime)
		return
	}
	if f.DayMaxNum == 0 || f.DayMaxNum >= DAYMAXPUBNUMBER {
		err = ecode.RequestErr
		log.Error("DayMaxNum(%v) is not validate", f.DayMaxNum)
		return
	}
	if f.PubTime == "" {
		err = ecode.RequestErr
		log.Error("PubTime(%v) is not validate", f.PubTime)
		return
	}
	if f.IsBatch == ISNOTBATCH {
		if f.Content == "" {
			err = ecode.RequestErr
			log.Error("Content(%v) is not validate", f.Content)
			return
		}
	}
	return
}

// sel form
type SelBusinessForm struct {
	Pn int64
	Ps int64
}

func (f *SelBusinessForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Pn: "pn",
		&f.Ps: "ps",
	}
}

func (f *SelBusinessForm) Validate(req *http.Request) (err error) {
	if f.Pn < 1 {
		err = ecode.RequestErr
		log.Error("Pn(%v) is not validate", f.Pn)
		return
	}
	if f.Ps < 1 {
		err = ecode.RequestErr
		log.Error("Ps(%v) is not validate", f.Ps)
		return
	}
	return
}
