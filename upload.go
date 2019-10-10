package main

import (
	"cartoon-gin/common"
	"cartoon-gin/myaws"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"io/ioutil"
	"log"
	"qiniupkg.com/x/bytes.v7"
	"strings"
)

const BASE_PATH = "/home/cartoon-spider/waiwai/"
//本地图片地址
var imageLocalPaths []string
//上传图片地址
var qiniuPaths []string
//封面地址 封面id ==> 地址
var chapterHoverImages map[string]string

var uploader *s3manager.Uploader
func UploadToAwsFromLocalFile()  {
	//上传aws
	sess, err := myaws.GetAwsSession()
	common.CheckError(err)
	uploader = s3manager.NewUploader(sess)


	//1.read image path
	bookId := "1194"
	path := BASE_PATH+bookId
	readImages(path)

	//2.upload
	for _,s := range imageLocalPaths {
		beginUpload(s)
	}
}

func readImages(path string) {
	if !common.DirExists(path) {
		log.Fatal("dir not exists:",path)
	}
	chapterDirs,err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("read dir error!path:"+path)
	}


	for _, dir := range chapterDirs {
		if dir.Name() == ".DS_Store" {
			continue
		}
		if dir.IsDir() {
			readImages(path+"/"+dir.Name())
		}else {
			//漫画封面
			p := path + "/" + dir.Name()
			imageLocalPaths = append(imageLocalPaths,p)
		}
	}
}

func beginUpload(localpath string) {
	str := strings.ReplaceAll(localpath,BASE_PATH,"")
	dirArr := strings.Split(str,"/")

	uploadPath := ""
	for _,v := range dirArr {
		if !strings.Contains(v,".jpg") {
			md5str := common.Md5Str(v)
			uploadPath += md5str[0:6]+"/"
		}else {
			//章节封面
			if strings.Contains(v, "cover") {

			}
			l := strings.Split(v,".")
			name := common.Md5Str(l[0])[0:8]
			uploadPath += name+".data"
		}
	}

	filename := uploadPath
	buffer := common.GetFileBufferFromLocal(localpath)
	//去掉前两个byte
	buffer = buffer[2:]
	reader := bytes.NewReader(buffer)
	toS3(filename,reader)
}

func toS3(filename string,fileBody io.Reader) {
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(myaws.S3_BUCKET),
		Key:         aws.String(filename),
		Body:        fileBody,
		ContentType: aws.String("image/jpeg"),
		ACL:         aws.String(s3.ObjectCannedACLPublicRead),
	})
	common.CheckError(err)

	fmt.Println("success -- ",filename)
}

