package dao

import "cartoon-gin/DB"

type Image struct {
	MyGormModel
	CartoonId int `json:"cartoon_id"`
	ChapterId int `json:"chapter_id"`
	ImageAddr string `json:"image_addr"`
	IsDeleted int `json:"is_deleted"`
}

func FindImageByChapterId(chapterId int) []Image {
	db,_ := DB.OpenCartoon()
	var list []Image
	db.Table("cartoon_images").Where("chapter_id = ?",chapterId).Scan(&list)
	return list
}