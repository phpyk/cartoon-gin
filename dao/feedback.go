package dao

import (
	"time"

	"cartoon-gin/DB"
)

type UserFeedback struct {
	ID           int       `gorm:"primary_key" json:"id"`
	CommitUserId int       `json:"commit_user_id"`
	Content      string    `json:"content"`
	Type         int       `json:"type"`
	ReplyId      int       `json:"reply_id"`
	CreatedAt    time.Time `json:"created_at" format:"2006-01-01 15:04:05"`
	UpdatedAt    time.Time `json:"-" format:"2006-01-01 15:04:05"`
}

const (
	//用户反馈
	FeedbackTypeCommit = 1
	//回复
	FeedbackTypeReply = 2
)

func (f *UserFeedback) Save() {
	db, _ := DB.OpenCartoon()
	if db.NewRecord(f) {
		db.Create(&f)
	}
}

//TODO 展示顺序有bug
func GetFeedbacks(userId, page, perPage int) []UserFeedback {
	db, _ := DB.OpenCartoon()
	var list []UserFeedback
	db.Table("user_feedbacks").Where("commit_user_id = ?",userId).
		Order("created_at DESC").
		Limit(perPage).
		Offset((page - 1) * perPage).
		Scan(&list)
	return list
}

func GetFeedbackCount(userId int) int {
	db, _ := DB.OpenCartoon()
	var count int
	db.Table("user_feedbacks").Where("commit_user_id = ?",userId).Count(&count)
	return count
}