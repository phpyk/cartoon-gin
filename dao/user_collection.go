package dao

import (
	"time"

	"cartoon-gin/DB"
	"github.com/jinzhu/gorm"
)

type UserCollection struct {
	MyGormModel
	UserId    int `json:"user_id"`
	CartoonId int `json:"cartoon_id"`
}

type BookCaseSearchRequest struct {
	Tab         string `form:"tab" json:"tab"`
	UserId      int    `form:"user_id" json:"user_id"`
	SortType    int    `form:"sort_type" json:"sort_type"`
	Page        int    `form:"page" json:"page"`
	PerPage     int    `form:"per_page" default:"18" json:"per_page"`
	IsAndroid   bool   `json:"is_android"`
	IsVerifying bool   `json:"is_verifying"`
	ShowRated   bool   `json:"show_rated"`
}

func GetUserCollections(searchRequest BookCaseSearchRequest) []map[string]interface{} {
	query := generalUserCollectionQuery(searchRequest)

	var list []QueryCartoons
	query.Limit(searchRequest.PerPage).
		Offset((searchRequest.Page - 1) * searchRequest.PerPage).
		Find(&list)
	return formatQueryCartoons(list)
}

func GetUserCollectionCount(searchRequest BookCaseSearchRequest) int {
	query := generalUserCollectionQuery(searchRequest)
	var totalCount int
	query.Count(&totalCount)
	return totalCount
}

func CartoonHasBeenCollected(userId, cartoonId int) bool {
	db, _ := DB.OpenCartoon()
	var count int
	db.Table("user_collections").Where("user_id = ?", userId).Where("cartoon_id = ?", cartoonId).Count(&count)
	return count > 0
}

func CollectCartoon(userId, cartoonId int) bool {
	record := UserCollection{UserId: userId, CartoonId: cartoonId}
	db, _ := DB.OpenCartoon()
	if db.NewRecord(record) {
		db.Create(&record)
		return true
	} else {
		return false
	}
}

func CancelCollectCartoon(userId, cartoonId int) bool {
	db, _ := DB.OpenCartoon()
	var record UserCollection
	db.Table("user_collections").Where("user_id = ?", userId).Where("cartoon_id = ?", cartoonId).First(&record)
	if record.ID > 0 {
		db.Delete(&record)
		return true
	}
	return false
}

func DeleteUserCollection(UserId int, cartoonIdSlice []int) bool {
	db, _ := DB.OpenCartoon()
	db.Table("user_collections").Where("user_id = ?", UserId).Where("cartoon_id in (?)", cartoonIdSlice).Update("deleted_at", time.Now().Format("2006-01-02 15:04:05"))
	return true
}

func generalUserCollectionQuery(searchRequest BookCaseSearchRequest) *gorm.DB {
	db, _ := DB.OpenCartoon()
	columns := "c.*"
	query := db.Debug().Table("user_collections AS a").
		Select(columns).
		Joins("INNER JOIN cartoons AS c ON a.cartoon_id = c.id").
		Where("c.is_on_sale = ?", CartoonIsOnSale).
		Where("c.verify_status = ?", CartoonVerifyStatusPass).
		Where("a.user_id = ?", searchRequest.UserId).
		Where("c.deleted_at IS NULL").
		Where("a.deleted_at IS NULL")
	if searchRequest.IsAndroid && searchRequest.IsVerifying {
		query = query.Where("c.cartoon_type = ?", CartoonTypeExternal)
	} else {
		query = query.Where("c.cartoon_type != ?", CartoonTypeExternal)
	}
	if !searchRequest.ShowRated {
		query = query.Where("c.is_rated = ?", 0)
	}
	if searchRequest.SortType == 3 {
		query = query.Order("a.updated_at DESC")
	} else if searchRequest.SortType == 2 {
		query = query.Order("c.updated_at DESC")
	} else {
		query = query.Order("a.created_at ASC")
	}
	return query
}
