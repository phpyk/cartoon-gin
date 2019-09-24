package main

import (
	"cartoon-gin/common"
	"cartoon-gin/myaws"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {

	//初始化session
	sess,err := myaws.GetAwsSession()
	common.CheckError(err)

	uploader := s3manager.NewUploader(sess)

	src := "http://cartoon.qiniu.tblinker.com/d972bc/331abc/08324b.jpg"
	//filename := myaws.ReadSrcAndLocalSave(src)
	filename,filebody := myaws.GetFileBodyAndName(src)
	//filebody,err := os.Open(filename)
	common.CheckError(err)
	_,err = uploader.Upload(&s3manager.UploadInput{
		Bucket:aws.String(myaws.S3_BUCKET),
		Key:aws.String(filename),
		Body:filebody,
		ContentType:aws.String("image/jpeg"),
		ACL: aws.String(s3.ObjectCannedACLPublicRead),
		//GrantFullControl:aws.String("FULL_CONTROL"),
		//GrantRead:aws.String(storagegateway.ObjectACLPublicRead),
	})

	common.CheckError(err)
	fmt.Printf("Successfully uploaded %q to %q \n", filename,  myaws.S3_BUCKET)
}