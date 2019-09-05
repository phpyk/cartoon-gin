package common

import (
	"crypto/md5"
	"regexp"
)

func IsPhone(phone string) bool {
	match,_ := regexp.MatchString(`^1[3456789]\d{9}$`,phone)
	return match
}

func encryPwd(originPwd string) string {
	enPwd := md5.New()
	enPwd.Sum()
}