package models

import "time"

type Book struct {
	ID        uint `gorm:"primary_key"`
	Name string
	Author string `gorm:"size:32"`
	Desc string
	WordCount float32
	Kinds string
	CoverUrl string `gorm:"size:500"`
	Source int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
