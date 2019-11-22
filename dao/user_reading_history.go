package dao

type UserReadingHistory struct {
	MyGormModel
	UserId int `json:"user_id"`
	CartoonId int `json:"cartoon_id"`
	ChapterId int `json:"chapter_id"`
	LastReadTime int `json:"last_read_time" gorm:"defa"`
}

func Add(userId,CartoonId,ChapterId int) bool {

}