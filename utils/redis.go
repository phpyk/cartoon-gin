package utils

import (
	"strings"

	"cartoon-gin/configs"
	"github.com/go-redis/redis"
)

/**
 * 全局使用的key，后面不要加冒号
 * 需要拼接其他参数的key，后面加冒号
 */
const (
	//手机验证码
	RDS_KEY_SMS_CODE string = "vcode:{phone}"
	//用户上次阅读章节
	RDS_KEY_LAST_READ_CHAPTER = "lastread:u:{uid}:c:{cid}"
	//Apple Pay用户，用来判断是否是审核人员
	RDS_KEY_APPLE_PAY_USERS = "apple_pay_users"

	RDS_KEY_USER_HOME_PAGE_RECOMMEND_DATA = "user:home:recomm:{uid}"

	RDS_KEY_USER_MORE_RECOMMEND_DATA = "user:more:recomm:{uid}:{p}"

	//用户已购章节
	RDS_KEY_USER_BUY_CHAPTERS = "user:buy:chapters:{uid}"
	//用户已读历史记录
	RDS_KEY_USER_READ_HISTORIES = "user:read:histories:{uid}"
	//用户推荐：漫画详情页
	RDS_KEY_USER_CARTOON_DETAIL_RECOMMEND_DATA = "user:detail:recomm:{uid}:{c}"
	//签到配置
	RDS_KEY_SIGN_IN_CONFIGS = "sign:in:configs"
)

func GetRedisKey(originalKey string, params map[string]string) string {
	var formatKey string = originalKey
	for k,v := range params {
		formatKey = strings.Replace(formatKey,"{"+k+"}",v,-1)
	}
	return formatKey
}

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     configs.REDIS_HOST + ":" + configs.REDIS_PORT,
		Password: configs.REDIS_PASSWORD,
		DB:       configs.REDIS_DB_DEFAULT,
	})
}

func NewAuthRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     configs.REDIS_HOST + ":" + configs.REDIS_PORT,
		Password: configs.REDIS_PASSWORD,
		DB:       configs.REDIS_DB_AUTHORIZATION,
	})
}

func RedisSave(k,v string) *redis.StatusCmd {
	clt := NewRedisClient()
	return clt.Set(k,v,0)
}