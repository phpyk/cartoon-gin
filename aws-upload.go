package main

import (
	"cartoon-gin/common"
	"cartoon-gin/dao"
	"cartoon-gin/myaws"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	//初始化session
	sess, err := myaws.GetAwsSession()
	common.CheckError(err)

	uploader := s3manager.NewUploader(sess)

	cpid := 4196
	imgList := dao.FindImagesByCartoonId(cpid)

	for _, row := range imgList {
		src := row.ImageAddr
		filename, filebody := myaws.GetFileBodyAndName(src)
		common.CheckError(err)

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket:      aws.String(myaws.S3_BUCKET),
			Key:         aws.String(filename),
			Body:        filebody,
			ContentType: aws.String("image/jpeg"),
			ACL:         aws.String(s3.ObjectCannedACLPublicRead),
		})

		common.CheckError(err)
		fmt.Printf("Successfully uploaded %q to %q \n", filename, myaws.S3_BUCKET)
	}

}
