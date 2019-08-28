package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Name string
	Author string `gorm:"size:32"`
	Score float64
	DescShort string
	DescLong string `sql:"type:text"`
	WordCount float64
	Kinds string
	CoverUrl string `gorm:"size:500"`
	Source int
	IsEnd int
	OriginBookId int
	BookDetailUrl string
}

func (b Book) TableName() string {
	return "books"
}
