package models

import "github.com/jinzhu/gorm"

type Keywords struct {
	gorm.Model
	Words string
}

func (k Keywords) TableName() string {
	return "keywords"
}