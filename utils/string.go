package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
)

func IsPhone(phone string) bool {
	match, _ := regexp.MatchString(`^1[3456789]\d{9}$`, phone)
	return match
}

func Md5Str(originPwd string) string {
	e1 := md5.New()
	e1.Write([]byte(originPwd))
	enstr := e1.Sum(nil)
	return hex.EncodeToString(enstr)
}

func EncryptPwd(originPwd string) string {
	e1 := md5.New()
	e1.Write([]byte(originPwd))
	b1 := e1.Sum(nil)

	e2 := md5.New()
	e2.Write([]byte(hex.EncodeToString(b1)))
	b2 := e2.Sum(nil)
	return hex.EncodeToString(b2)
}

func RandomString(l, strType int) string {
	var c string
	if strType == 1 {
		c = "0123456789"
	} else if strType == 2 {
		c = "abcdefghijklmnopqrstuvwxyz"
	} else if strType == 3 {
		c = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if strType == 4 {
		c = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	} else if strType == 5 {
		c = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	} else {
		c = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	b := make([]byte, l)
	for i := 0; i < l; i++ {
		idx := rand.Intn(len(c))
		b[i] = c[idx]
	}
	return string(b)
}

func GeneralInviteCode() string {
	code := RandomString(6, 5)
	return code
}

func GeneralNickName() string {
	pre := RandomString(6, 4)
	end := RandomString(4, 1)
	return pre + end
}
