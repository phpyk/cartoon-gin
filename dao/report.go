package dao

import (
	"cartoon-gin/DB"
	"github.com/jinzhu/gorm"
)

const (
	ReportReasonTypeQueshi    = 1
	ReportReasonTypeDisu      = 2
	ReportReasonTypeQinquan   = 3
	ReportReasonTypeHuazhicha = 4
	ReportReasonTypeOthers    = 5
)

type UserReport struct {
	MyGormModel
	UserId     int    `json:"user_id"`
	CartoonId  int    `json:"cartoon_id"`
	ReasonType int    `json:"reason_type"`
	Reason     string `json:"reason"`
}

func (r *UserReport) GetReportReason(reportType int) string {
	var reason string
	switch reportType {
	case ReportReasonTypeQueshi:
		reason = "内容缺失"
	case ReportReasonTypeDisu:
		reason = "内容低俗"
	case ReportReasonTypeQinquan:
		reason = "内容侵权"
	case ReportReasonTypeHuazhicha:
		reason = "画质差"
	default:
		reason = "其他"
	}
	return reason
}

func (r *UserReport) Save() *gorm.DB {
	db, _ := DB.OpenCartoon()
	if db.NewRecord(r) {
		return db.Save(&r)
	}
	return nil
}
