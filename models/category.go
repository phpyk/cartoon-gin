package models

import (
	"cartoon-gin/common"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	CateName string
	ParentId uint
}

func (c Category) TableName() string {
	return "categories"
}

func GetAllCategories() (cats []Category) {
	db,err := OpenNewsDB()
	common.CheckError(err)
	db.Debug().Find(&cats)
	return cats
}

func AddCategory(catName string, parentId uint) bool {
	db,err := OpenNewsDB()
	common.CheckError(err)
	category := Category{CateName:catName,ParentId:parentId}
	return db.NewRecord(category)
}