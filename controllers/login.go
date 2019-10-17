package controllers

import (
	"cartoon-gin/DB"
	"cartoon-gin/auth"
	"cartoon-gin/utils"
	"cartoon-gin/dao"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

//LoginAction handle login by phone and password
func LoginAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	phone := c.Request.FormValue("phone")
	password := c.Request.FormValue("password")

	if !utils.IsPhone(phone) {
		cg.Failed("手机号格式不正确")
	}
	user := dao.UserFindByPhone(phone)
	if !(user.ID > 0) {
		cg.Failed("用户不存在")
	}
	encryptPwd := utils.EncryptPwd(password)
	if encryptPwd != user.Password {
		cg.Failed("密码不正确")
	}

	resp, err := loginUser(user, c)
	if err != nil {
		cg.Failed("login faild via:" + err.Error())
	}
	cg.Success(resp)
}

//VisitorLoginAction create a visitor and login
func VisitorLoginAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	if _, ok := c.Request.Header["H-Device"]; !ok {
		cg.Failed("device_token is required")
	}
	devices := c.Request.Header["H-Device"]
	//replace blank
	device := strings.ReplaceAll(devices[0], " ", "")
	if len(device) == 0 {
		cg.Failed("device_token is required")
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
		user.UserType = dao.USER_TYPE_VISITOR
		user.UserDevice = device
		dao.UserCreate(user)
	}
	resp, err := loginUser(user, c)

	if err != nil {
		cg.Failed("login faild via:" + err.Error())
	}
	cg.Success(resp)
}

func LogoutAction(c *gin.Context) {

}

func CurrentUserAction(c *gin.Context) {
	cg := utils.Gin{C: c}
	//interface 转 uint类型
	//cg.C.Keys["uid"].(uint)
	me := dao.UserFindByID(cg.C.Keys["uid"].(uint))
	cg.Success(me)
}

//loginUser handle user login
func loginUser(user dao.User, c *gin.Context) (map[string]interface{}, error) {
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
