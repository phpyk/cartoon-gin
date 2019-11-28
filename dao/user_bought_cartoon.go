package dao

import (
	"fmt"

	"cartoon-gin/DB"
	"github.com/jinzhu/gorm"
)

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

func BuyChapter(user *User,chapter *Chapter) error {
	db,_ := DB.OpenCartoon()
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	//1.增加购买记录
	buyRecord := UserBoughtCartoon{UserId: user.ID, CartoonId: chapter.CartoonId, ChapterId: chapter.ID,}
	if tx.NewRecord(buyRecord) {
		if err := tx.Create(&buyRecord).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	//2.增加金币消费记录
	balance := user.ValidCoin - uint(chapter.SalePrice)
	coinRecord := CoinRecord{UserId: user.ID, Amount: chapter.SalePrice, Balance: int(balance), ActType: CoinActTypeDecrease, CartoonId: chapter.CartoonId, ChapterId: chapter.ID, Remark: "购买章节", ReferBizId: buyRecord.ID,}
	if tx.NewRecord(coinRecord) {
		if err := tx.Create(&coinRecord).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	//3.扣除金币
	if err := tx.Model(user).Update("valid_coin", balance).Error; err != nil {
		fmt.Println("扣除金币error：",err)
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func GetUserBoughtCartoons(searchRequest BookCaseSearchRequest) []map[string]interface{} {
	query := generalUserBoughtCartoonQuery(searchRequest)
	var list []QueryCartoons
	query.Limit(searchRequest.PerPage).
		Offset((searchRequest.Page - 1) * searchRequest.PerPage).
		Scan(&list)
	return formatQueryCartoons(list)
}

func GetUserBoughtCartoonsCount(searchRequest BookCaseSearchRequest) int {
	query := generalUserBoughtCartoonQuery(searchRequest)
	var count int
	query.Count(&count)
	return count
}

func generalUserBoughtCartoonQuery(searchRequest BookCaseSearchRequest) *gorm.DB {
	db,_ := DB.OpenCartoon()
	columns := "c.*"
	query := db.Debug().Select(columns).Table("user_bought_cartoons AS b").
		Joins("INNER JOIN cartoons AS c ON b.cartoon_id = c.id").
		Where("c.is_on_sale = ?",CartoonIsOnSale).
		Where("c.verify_status = ?",CartoonVerifyStatusPass).
		Where("b.user_id = ?",searchRequest.UserId).
		Where("c.deleted_at IS NULL").
		Where("b.deleted_at IS NULL")

	if !searchRequest.ShowRated {
		query = query.Where("c.is_rated = ?", 0)
	}
	query = query.Group("b.user_id, b.cartoon_id")
	return query
}