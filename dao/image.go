package dao

import "cartoon-gin/DB"

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

func FindImagesByCartoonId(cartoonId int) []Image {
	db, _ := DB.OpenCartoon()
	var list []Image
	db.Table("cartoon_images").Where("cartoon_id = ?", cartoonId).Scan(&list)
	return list
}

func FindImagesForUpload(offset int) []Image {
	db, _ := DB.OpenCartoon()
	var list []Image
	db.Table("cartoon_images").Limit(1000).Offset(offset).Scan(&list)
	return list
}
