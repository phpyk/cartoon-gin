package dao

import (
	"cartoon-gin/DB"
)

type Image struct {
	MyGormModel
	CartoonId int    `json:"cartoon_id"`
	ChapterId int    `json:"chapter_id"`
	ImageAddr string `json:"image_addr"`
	IsDeleted int    `json:"is_deleted"`
}

func FindImagesByChapterId(chapterId int) []Image {
	db, _ := DB.OpenCartoon()
	var list []Image
	db.Table("cartoon_images").Where("chapter_id = ?", chapterId).Scan(&list)
	return list
}

func FindImagesListById(ids []string) []Image {
	db, _ := DB.OpenCartoon()
	var list []Image
	db.Table("cartoon_images").Where("id in (?)", ids).Scan(&list)
	return list
}
func FindImagesByCartoonId(cartoonId int) []Image {
	db, _ := DB.OpenCartoon()
	var list []Image
	db.Table("cartoon_images").Where("cartoon_id = ?", cartoonId).Scan(&list)
	return list
}

func FindImagesForUpload(limit int,lastMaxId int) []Image {
	db, _ := DB.OpenCartoon()
	var list []Image
	db.Table("cartoon_images").Where("id > ?",lastMaxId).Limit(limit).Order("id ASC").Scan(&list)
	return list
}
