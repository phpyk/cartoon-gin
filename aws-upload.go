package main

import (
	"cartoon-gin/common"
	"cartoon-gin/dao"
	"cartoon-gin/myaws"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"qiniupkg.com/x/log.v7"
	"strconv"
	"time"
)

var idfile = "max_id.log"

func main() {
	limit := flag.Int("limit",1000,"limit")
	lastMaxId := getMaxId()

	timeBegin := time.Now()
	//初始化session
	sess, err := myaws.GetAwsSession()
	common.CheckError(err)

	uploader := s3manager.NewUploader(sess)

	imgList := dao.FindImagesForUpload(*limit,lastMaxId)

	fmt.Printf("count:%d \n",len(imgList))

	maxId := 0
	for _, row := range imgList {
		t1 := time.Now()
		src := row.ImageAddr
		filename, filebody := myaws.GetFileBodyAndName(src)

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket:      aws.String(myaws.S3_BUCKET),
			Key:         aws.String(filename),
			Body:        filebody,
			ContentType: aws.String("image/jpeg"),
			ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		})
		if err != nil {
			log.Println("error row_id=%v",row.ID)
			continue
		}

		escaped := time.Since(t1)
		fmt.Printf("Successfully uploaded %q to %q -- escaped: %v\n", filename, myaws.S3_BUCKET,escaped)
		if row.ID > maxId {
			maxId = row.ID
			saveMaxId(maxId)
		}
	}

	totalTime := time.Since(timeBegin)
	fmt.Printf("Done total_time:%v \n",totalTime)
}

func getMaxId() int {
	c,err := ioutil.ReadFile(idfile)
	common.CheckError(err)
	i,err := strconv.Atoi(string(c))
	common.CheckError(err)
	return i
}
func saveMaxId(id int) {
	b := []byte(strconv.Itoa(id))
	err := ioutil.WriteFile(idfile,b,0644)
	common.CheckError(err)
}
