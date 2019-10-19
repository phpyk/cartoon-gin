package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type SmsClient struct {
	AccountSid   string
	AccountToken string
	AppId        string
	ServerIp     string
	ServerPort   string
	Version      string
	Batch        string //时间戳
	BodyType     string //包体格式，可填值：json 、xml
	EnableLog    bool   //日志开关。可填值：true/false
	FileName     string //日志文件
}
type ClientError struct {
	Message string
	e       error
}

func NewSmsClient() SmsClient {
	return SmsClient{
		AccountSid:   "8aaf07086812057f01682687c85208e2",
		AccountToken: "ec44618e428146f79e4a7f203a759caa",
		AppId:        "8aaf07086a25761e016a2e38513d0552",
		ServerIp:     "app.cloopen.com",
		ServerPort:   "8883",
		Version:      "2013-12-26",
		BodyType:     "json",
		Batch:        time.Now().Format("20060102150405"),
	}
}

func (c *SmsClient) accountCheck() error {
	if c.ServerIp == "" {
		return errors.New("IP为空")
	}
	//port,err := strconv.Atoi(c.ServerPort)
	//if port <= 0 || err != nil {
	//	return errors.New("端口错误（小于等于0）")
	//}
	if c.Version == "" {
		return errors.New("版本号为空")
	}
	if c.AccountSid == "" {
		return errors.New("主帐号为空")
	}
	if c.AccountToken == "" {
		return errors.New("主帐号令牌为空")
	}
	if c.AppId == "" {
		return errors.New("应用ID为空")
	}
	return nil
}

func (c *SmsClient) SendTemplateSMS(to string, data []string, templateId int) (result map[string]interface{}, err error) {
	err = c.accountCheck()
	if err != nil {
		return nil, err
	}
	var datas string
	var body string
	templateIdStr := strconv.Itoa(templateId)
	// 拼接请求包体
	for _, v := range data {
		datas += "'" + v + "',"
	}
	body = "{'to':'" + to + "','templateId':'" + templateIdStr + "','appId':'" + c.AppId + "','datas':[" + datas + "]}"
	// 大写的sig参数
	sign := strings.ToUpper(Md5Str(c.AccountSid + c.AccountToken + c.Batch))
	// 生成请求URL
	url := "https://" + c.ServerIp + ":" + c.ServerPort + "/" + c.Version + "/Accounts/" + c.AccountSid + "/SMS/TemplateSMS?sig=" + sign
	// 生成授权：主帐户Id + 英文冒号 + 时间戳。
	s := c.AccountSid + ":" + c.Batch
	authen := base64.StdEncoding.EncodeToString([]byte(s))
	// 添加包头
	request, err := http.NewRequest("POST", url, strings.NewReader(body))
	request.Header.Add("Accept", "application/"+c.BodyType)
	request.Header.Add("Content-Type", "application/"+c.BodyType+";charset=utf-8")
	request.Header.Add("Authorization", authen)
	// 发送请求
	clt := http.Client{}
	response, err := clt.Do(request)
	log.Printf("response:%+v\n", response.Body)

	resultB, err := ioutil.ReadAll(response.Body)
	log.Printf("resultB:  %+v", string(resultB))
	err = json.Unmarshal(resultB, &result)

	if err != nil {
		return nil, err
	}
	return result, nil
}
