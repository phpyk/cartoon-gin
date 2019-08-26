package models

import (
	"cartoon-gin/common"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	CateName string
	ParentId int
}

func (c Category) TableName() string {
	return "categories"
}

func GetAllCategories() (cats []Category) {
	db,err := OpenNewsDB()
	common.CheckError(err)
	db.Find(&cats)
	return cats
}

func AddCategory(catName string, parentId int) bool {
	db,err := OpenNewsDB()
	common.CheckError(err)
	category := Category{CateName:catName,ParentId:parentId}
	db.Create(&category)
	return true
}

func GetCatByName(catName string) Category {
	db,err := OpenNewsDB()
	common.CheckError(err)
	var cat Category
	db.Where("cate_name = ?",catName).First(&cat)
	return cat
}

func UpdateCatById(catName string,id int) bool {
	db,err := OpenNewsDB()
	common.CheckError(err)
	var cat Category
	db.Model(&cat).Where("id = ?",id).Update("cate_name",catName)
	return true
}