package dao

import "cartoon-gin/DB"

type Category struct {
	MyGormModel
	CatName string `json:"cat_name"`
	IsValid int `json:"is_valid"`
	BackgroundImg string `json:"background_img"`
	IconSelect string `json:"icon_select"`
	IconUnselect string `json:"icon_unselect"`
}

func GetAllCategories() (cats []Category) {
	db, _ := DB.OpenCartoon()
	db.Table("categories").Where("deleted_at IS NULL").Scan(&cats)
	return cats
}