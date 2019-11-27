package dao

import (
	"time"

	"cartoon-gin/DB"
)

type UserReadingHistory struct {
	MyGormModel
	UserId int `json:"user_id"`
	CartoonId int `json:"cartoon_id"`
	ChapterId int `json:"chapter_id"`
	LastReadTime int64 `json:"last_read_time"`
}

func AddToReadingHistory(userId,CartoonId,ChapterId int) bool {
	db,_ := DB.OpenCartoon()
	history := UserReadingHistory{UserId:userId,CartoonId:CartoonId,ChapterId:ChapterId,LastReadTime:time.Now().Unix()}
	if db.Table("user_reading_historys").NewRecord(history) {
		db.Table("user_reading_historys").Create(&history)
		return true
	}
	return false
}