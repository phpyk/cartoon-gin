package myaws

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"cartoon-gin/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetAwsSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(REGION),
		Credentials: credentials.NewStaticCredentials(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ""),
		Endpoint:    aws.String(S3_ENDPOINT),
	})
	return sess, err
}

func ReadSrcAndLocalSave(imageUrl string) (filename string) {
	resp, err := http.Get(imageUrl)
	utils.CheckError(err)
	body, err := ioutil.ReadAll(resp.Body)
	utils.CheckError(err)

	//分隔imageUrl,创建文件夹
	arr := strings.Split(imageUrl, "/")
	l := len(arr)
	dir1 := arr[l-3]
	utils.CreateDirIfNotExists(dir1)

	dir2 := dir1 + "/" + arr[l-2]
	utils.CreateDirIfNotExists(dir2)

	localFileName := dir2 + "/" + arr[l-1]
	utils.CreateFile(localFileName, body)
	return localFileName
}

func GetFileBodyAndName(imageHttpUrl string) (string, io.Reader) {
	resp, err := http.Get(imageHttpUrl)
	utils.CheckError(err)

	//分隔imageUrl
	idx := strings.Index(imageHttpUrl,".com/")

	substr := imageHttpUrl[idx+5:]
	arr := strings.Split(substr, "/")

	var name string
	for _,v := range arr {
		name += "/"+v
	}
	name = name[1:]
	return name, resp.Body
}

