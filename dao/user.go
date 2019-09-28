package dao

import (
	"cartoon-gin/DB"
)

const USER_TYPE_VISITOR = 0
const USER_TYPE_NORMAL = 1
const USER_TYPE_WILLING = 2
const USER_TYPE_TARGET = 3

type User struct {
	MyGormModel
	Phone                                                                                 string `json:"phone"`
	Password                                                                              string `json:"password"`
	NickName                                                                              string `json:"nick_name"`
	UserType                                                                              uint   `json:"user_type"`
	ValidCoin                                                                             uint   `json:"valid_coin"`
	IsVip, VipExpireTime, IsAutoRenewal, LastLoginTime, InviteUid, InviteTime, SignInDays uint
	LastLoginIp, UserDevice, InviteCode                                                   string
}

type JwtToken struct {
	Token string `json:"token"`
}

func UserFindByID(id uint) (user User) {
	db, _ := DB.OpenCartoon()
	db.Where("id = ?", id).First(&user)
	return user
}

func UserFindByPhone(phone string) (user User) {
	db, _ := DB.OpenCartoon()
	db.Where("phone = ?", phone).First(&user)
	return user
}
func UserFindByDeviceToken(devicetoken string) (user User) {
	db, _ := DB.OpenCartoon()
	db.Where("user_device = ?",devicetoken).First(&user)
	return user
}

func UserInviteCodeExists(code string) bool {
	db, _ := DB.OpenCartoon()
	c := 0
	db.Table("users").Where("invite_code = ?",code).Count(&c)
	return c > 0
}
func UserCreate(user User) bool {
	db, _ := DB.OpenCartoon()
	db.NewRecord(user)
	db.Create(&user)
	db.First(&user)
	return user.ID > 0
}

