package models

import "github.com/jinzhu/gorm"

type Chapter struct {
	gorm.Model
	Name string `gorm:"size:64"`
	WordCount int
	Content string `sql:"type:text"`
}
