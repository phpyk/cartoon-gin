package dao

import (
	"cartoon-gin/common"
	"time"
)

type MyGormModel struct {
	ID        uint               `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time      `json:"updated_at" time_format:"2006-01-02 15:04:05"`
	DeletedAt *common.MyNullTime `json:"deleted_at" `
}

type Paginate struct {
	Page int `gorm:"default:1" json:"page"`
	PerPage int `gorm:"default:20" json:"per_page"`
}