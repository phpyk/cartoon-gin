package common

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

func IsPhone(phone string) bool {
	match,_ := regexp.MatchString(`^1[3456789]\d{9}$`,phone)
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