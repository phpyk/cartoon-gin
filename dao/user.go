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
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	NickName      string `json:"nick_name"`
	UserType      uint   `json:"user_type"`
	ValidCoin     uint   `json:"valid_coin"`
	IsVip         uint   `json:"is_vip"`
	VipExpireTime uint   `json:"vip_expire_time"`
	IsAutoRenewal uint   `json:"is_auto_renewal"`
	LastLoginTime uint   `json:"last_login_time"`
	InviteUid     uint   `json:"invite_uid"`
	InviteTime    uint   `json:"invite_time"`
	SignInDays    uint   `json:"sign_in_days"`
	LastLoginIp   string `json:"last_login_ip"`
	UserDevice    string `json:"user_device"`
	InviteCode    string `json:"invite_code"`
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
	db.Where("user_device = ?", devicetoken).First(&user)
	return user
}

func UserInviteCodeExists(code string) bool {
	db, _ := DB.OpenCartoon()
	c := 0
	db.Table("users").Where("invite_code = ?", code).Count(&c)
	return c > 0
}
func UserCreate(user User) bool {
	db, _ := DB.OpenCartoon()
	db.NewRecord(user)
	db.Create(&user)
	db.First(&user)
	return user.ID > 0
}
