package controllers

import (
	"strings"
	"time"

	"cartoon-gin/DB"
	"cartoon-gin/auth"
	"cartoon-gin/dao"
	"cartoon-gin/utils"
	"github.com/gin-gonic/gin"
)

//PasswordLoginAction handle login by phone and password
func PasswordLoginAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	phone := c.Request.FormValue("phone")
	password := c.Request.FormValue("password")

	if !utils.IsPhone(phone) {
		cg.Failed("手机号格式不正确")
		return
	}
	user := dao.UserFindByPhone(phone)
	if !(user.ID > 0) {
		cg.Failed("用户不存在")
		return
	}
	encryptPwd := utils.EncryptPwd(password)
	if encryptPwd != user.Password {
		cg.Failed("密码不正确")
		return
	}

	//resp, err := loginUserWithJWT(user, c)
	resp, err := loginUserWithRedisAuth(user, c)
	if err != nil {
		cg.Error("login faild via:" + err.Error())
		return
	}
	cg.Success(resp)
}

func LoginAction(c *gin.Context) {
	cg := utils.Gin{C: c,}
	phone := c.Request.FormValue("phone")
	if !utils.IsPhone(phone) {
		cg.Failed("手机号格式不正确")
		return
	}
	smsCode := c.Request.FormValue("verify_code")
	if !checkSmsCode(phone,smsCode) {
		cg.Failed("手机号验证码不正确")
		return
	}
	user := dao.UserFindByPhone(phone)
	if !(user.ID > 0) {
		cg.Failed("用户不存在")
		return
	}
	//resp, err := loginUserWithJWT(user, c)
	resp, err := loginUserWithRedisAuth(user, c)
	if err != nil {
		cg.Error("login faild via:" + err.Error())
		return
	}
	cg.Success(resp)
}

//VisitorLoginAction create a visitor and login
func VisitorLoginAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	if _, ok := c.Request.Header["H-Device"]; !ok {
		cg.Failed("device_token is required")
		return
	}
	devices := c.Request.Header["H-Device"]
	//replace blank
	device := strings.ReplaceAll(devices[0], " ", "")
	if len(device) == 0 {
		cg.Failed("device_token is required")
		return
	}
	user := dao.UserFindByDeviceToken(device)
	//create a visitor user
	if user.ID == 0 {
		var code string
		code = utils.GeneralInviteCode()
		for dao.UserInviteCodeExists(code) {
			code = utils.GeneralInviteCode()
		}
		user.InviteCode = code
		user.NickName = utils.GeneralNickName()
		user.UserType = dao.UserTypeVisitor
		user.UserDevice = device
		dao.UserCreate(&user)
	}
	//resp, err := loginUserWithJWT(user, c)
	resp, err := loginUserWithRedisAuth(user, c)

	if err != nil {
		cg.Error("login faild via:" + err.Error())
		return
	}
	cg.Success(resp)
}

func LoginDevAccount(c *gin.Context) {
	cg := utils.Gin{C:c}
	user := dao.UserFindByPhone("17505818455")
	if !(user.ID > 0) {
		cg.Failed("用户不存在")
		return
	}
	//resp, err := loginUserWithJWT(user, c)

	resp,err  := loginUserWithRedisAuth(user,c)
	if err != nil {
		cg.Error("login faild via:" + err.Error())
		return
	}
	cg.Success(resp)
}

func LogoutAction(c *gin.Context) {
	cg := utils.Gin{C: c}

	token := c.Request.Header.Get("Authorization")
	if token == "" || len(token) <= 7 {
		cg.UnAuthorized()
	}
	token = token[7:]
	clt := utils.NewAuthRedisClient()
	sts := clt.Del(token)
	_,err := sts.Result()
	if err != nil {
		cg.Failed("退出失败")
	}
	cg.Success(nil)
}

func CurrentUserAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	//interface 转 uint类型
	//cg.C.Keys["uid"].(uint)
	user := cg.C.Keys["user"]
	if user == nil {
		cg.Failed("用户未登陆")
		return
	}
	//me := dao.UserFindByID(cg.C.Keys["uid"].(uint))
	me := cg.C.Keys["user"].(dao.User)
	cg.Success(me)
}

//loginUserWithJWT handle user login
func loginUserWithJWT(user dao.User, c *gin.Context) (map[string]interface{}, error) {
	lastLoginTime := uint(time.Now().Unix())
	lastLoginIp := c.Request.RemoteAddr
	db, _ := DB.OpenCartoon()
	db.Model(&user).Updates(dao.User{LastLoginIp: lastLoginIp, LastLoginTime: lastLoginTime})

	token, err := auth.GenerateToken(&user)

	resp := make(map[string]interface{})
	resp["token"] = token
	resp["token_type"] = "Bearer"
	resp["user_info"] = user
	return resp, err
}
func loginUserWithRedisAuth(user dao.User,c *gin.Context) (map[string]interface{}, error) {
	lastLoginTime := uint(time.Now().Unix())
	lastLoginIp := c.Request.RemoteAddr
	db, _ := DB.OpenCartoon()
	db.Model(&user).Updates(dao.User{LastLoginIp: lastLoginIp, LastLoginTime: lastLoginTime})

	token,err := auth.GenerateRedisToken(&user)

	resp := make(map[string]interface{})
	resp["token"] = token
	resp["token_type"] = ""
	resp["user_info"] = user
	return resp, err
}

func checkSmsCode(phone, code string) bool {
	redisClient := utils.NewRedisClient()
	par := make(map[string]string)
	par["phone"] = phone
	key := utils.GetRedisKey(utils.RDS_KEY_SMS_CODE,par)
	redisCode,err := redisClient.Get(key).Result()
	utils.CheckError(err)
	return redisCode == code
}

