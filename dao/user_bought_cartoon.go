package dao

import "cartoon-gin/DB"

type UserBoughtCartoon struct {
	MyGormModel
	UserId int `json:"user_id"`
	CartoonId int `json:"cartoon_id"`
	ChapterId int `json:"chapter_id"`
}

func HasBoughtChapter(userId, chapterId int) bool {
	var c int
	db,_ := DB.OpenCartoon()
	db.Table("user_bought_cartoons").Where("user_id = ?", userId).Where("chapter_id = ?",chapterId).Count(&c)
	return c > 0
}