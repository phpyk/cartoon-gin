package models

import "time"

type Book struct {
	ID        uint `gorm:"primary_key"`
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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
