package main

import (
	"fmt"
	"os"

	"cartoon-gin/common"
	"cartoon-gin/myaws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {

	//初始化session
	sess,err := myaws.GetAwsSession()
	common.CheckError(err)

	uploader := s3manager.NewUploader(sess)

	src := "http://cartoon1.qiniu.tblinker.com/30c8e1/e52e41/45c48cce.data"
	filename := myaws.ReadSrcAndLocalSave(src)

	file,err := os.Open(filename)
	common.CheckError(err)
	_,err = uploader.Upload(&s3manager.UploadInput{
		Bucket:aws.String(myaws.S3_BUCKET),
		Key:aws.String(filename),
		Body:file,
	})

	if err != nil {
		// Print the error and exit.
		fmt.Printf("Unable to upload %q to %q, %v \n", filename, myaws.S3_BUCKET, err)
	}else {
		fmt.Printf("Successfully uploaded %q to %q \n", filename,  myaws.S3_BUCKET)
	}
}