package dao

import (
	"cartoon-gin/DB"
)

type User struct {
	MyGormModel
	Phone string `json:"phone"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	UserType uint `json:"user_type"`
	ValidCoin uint	`json:"valid_coin"`
	IsVip,VipExpireTime,IsAutoRenewal,LastLoginTime,InviteUid,InviteTime,SignInDays uint
	LastLoginIp,UserDevice,InviteCode string
}

func FindUserByPhone(phone string) (user User) {
	db,_ := DB.OpenCartoon()
	db.Where("phone = ?",phone).First(&user)
	return user
}
