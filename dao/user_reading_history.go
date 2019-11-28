package dao

import (
	"time"

	"cartoon-gin/DB"
	"github.com/jinzhu/gorm"
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

func GetUserReadingHistories(searchRequest BookCaseSearchRequest) []map[string]interface{} {
	query := generalUserReadingHistoryQuery(searchRequest)
	var list []QueryCartoons
	query.Order("h.last_read_time DESC").
		Limit(searchRequest.PerPage).
		Offset((searchRequest.Page - 1) * searchRequest.PerPage).
		Scan(&list)
	return formatQueryCartoons(list)
}

func GetUserReadingHistoryCount(searchRequest BookCaseSearchRequest) int {
	query := generalUserReadingHistoryQuery(searchRequest)
	var count int
	query.Count(&count)
	return count
}

func DeleteUserReadingHistory(UserId int, cartoonIdSlice []int) bool {
	db,_ := DB.OpenCartoon()
	db.Table("user_reading_historys").Where("user_id = ?",UserId).Where("cartoon_id in (?)",cartoonIdSlice).Update("deleted_at",time.Now().Format("2006-01-02 15:04:05"))
	return true
}

func generalUserReadingHistoryQuery(searchRequest BookCaseSearchRequest) *gorm.DB {
	db,_ := DB.OpenCartoon()
	columns := "h.cartoon_id,c.*, max(h.chapter_id) as last_read_chapter_id, max(h.last_read_time) as last_read_time"
	query := db.Debug().Table("user_reading_historys AS h").
		Select(columns).
		Joins("INNER JOIN cartoons AS c ON h.cartoon_id = c.id").
		Where("c.is_on_sale = ?", CartoonIsOnSale).
		Where("c.verify_status = ?",CartoonVerifyStatusPass).
		Where("c.cartoon_type != ?",CartoonTypeExternal).
		Where("h.user_id = ?",searchRequest.UserId).
		Where("c.deleted_at is null").
		Where("h.deleted_at IS NULL")

	if !searchRequest.ShowRated {
		query = query.Where("c.is_rated = ?", 0)
	}
	query = query.Group("h.cartoon_id")
	return query
}