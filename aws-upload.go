package main

import (
	"cartoon-gin/common"
	"cartoon-gin/myaws"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func upload() {
	tagName1 := "Cost Center"
	tagValue1 := "123456"
	tagName2 := "Stack"
	tagValue2 := "MyTestStack"

	//初始化session
	sess,err := session.NewSession(&aws.Config{
		Region: aws.String(myaws.REGION),
		Credentials: credentials.NewStaticCredentials(myaws.ACCESS_KEY_ID,myaws.SECRET_ACCESS_KEY,""),
		Endpoint:aws.String(myaws.S3_ENDPOINT),
	})
	common.CheckError(err)

	//设置log level
	svc := s3.New(sess,aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))

	pubInput := &s3.PutBucketTaggingInput{
		Bucket:aws.String(myaws.S3_BUCKET),
		Tagging:&s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String(tagName1),
					Value: aws.String(tagValue1),
				},
				{
					Key:   aws.String(tagName2),
					Value: aws.String(tagValue2),
				},
			},
		},
	}
	_,err = svc.PutBucketTagging(pubInput)
	common.CheckError(err)

	getInput := &s3.GetBucketTaggingInput{
		Bucket: aws.String(myaws.S3_BUCKET),
	}
	result,err := svc.GetBucketTagging(getInput)
	common.CheckError(err)

	numTags := len(result.TagSet)
	if numTags > 0 {
		fmt.Println("Found", numTags, "Tag(s):")
		fmt.Println("")

		for _, t := range result.TagSet {
			fmt.Println("  Key:  ", *t.Key)
			fmt.Println("  Value:", *t.Value)
			fmt.Println("")
		}
	} else {
		fmt.Println("Did not find any tags")
	}
}