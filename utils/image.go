package utils

import (
	"encoding/base64"
	"io/ioutil"
)

//ImageBase64String 获取图片的base64字符串，前面加图片类型
func ImageBase64String(filename string) string {
	filebyte, err := ioutil.ReadFile(filename)
	CheckError(err)

	strContent := base64.StdEncoding.EncodeToString(filebyte)
	prepend := "data:image/png;base64"
	return prepend + "," + strContent
}
