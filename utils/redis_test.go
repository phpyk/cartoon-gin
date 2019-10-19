package utils

import (
	"fmt"
	"testing"
)

func TestGetRedisKey(t *testing.T) {
	params := make(map[string]string)
	params["uid"] = "713"
	params["cid"] = "111"
	key := GetRedisKey(RDS_KEY_LAST_READ_CHAPTER,params)
	fmt.Println(key)
}

func TestRedisSave(t *testing.T) {
	params := make(map[string]string)
	params["phone"] = "17505818455"

	k := GetRedisKey(RDS_KEY_SMS_CODE,params)
	sts := RedisSave(k,"123123")
	res,err := sts.Result()
	CheckError(err)
	fmt.Printf("%+v\n",res)
}
