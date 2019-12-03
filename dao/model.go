package dao

import (
	"cartoon-gin/utils"
	"time"
)

type MyGormModel struct {
	ID        int               `gorm:"primary_key" json:"id"`
	CreatedAt time.Time         `json:"-" format:"2006-01-01 15:04:05"`
	UpdatedAt time.Time         `json:"-" format:"2006-01-01 15:04:05"`
	DeletedAt *utils.MyNullTime `json:"-"`
}

type Paginate struct {
	Page    int `gorm:"default:1" json:"page"`
	PerPage int `gorm:"default:18" json:"per_page"`
}
