package main

import (
	"cartoon-gin/common"
	"cartoon-gin/dao"
	"cartoon-gin/myaws"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"time"
)

func main() {
	timeBegin := time.Now()
	//初始化session
	sess, err := myaws.GetAwsSession()
	common.CheckError(err)

	uploader := s3manager.NewUploader(sess)

	imgList := dao.FindImagesForUpload(0)
	fmt.Printf("count:%d \n",len(imgList))

	for _, row := range imgList {
		t1 := time.Now()
		src := row.ImageAddr
		filename, filebody := myaws.GetFileBodyAndName(src)
		fmt.Println(filename)
		common.CheckError(err)

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket:      aws.String(myaws.S3_BUCKET),
			Key:         aws.String(filename),
			Body:        filebody,
			ContentType: aws.String("image/jpeg"),
			ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		})

		common.CheckError(err)

		escaped := time.Since(t1)
		fmt.Printf("Successfully uploaded %q to %q -- escaped: %v\n", filename, myaws.S3_BUCKET,escaped)
	}

	totalTime := time.Since(timeBegin)
	fmt.Printf("Done total_time:%d \n",totalTime)
}
