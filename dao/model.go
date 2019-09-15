package dao

import (
	"cartoon-gin/common"
)

type MyGormModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt common.MyTime `json:"created_at"`
	UpdatedAt common.MyTime `json:"updated_at"`
	DeletedAt *common.MyNullTime `json:"deleted_at"`
}
