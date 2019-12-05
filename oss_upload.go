package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"cartoon-gin/configs"
	"cartoon-gin/dao"
	"cartoon-gin/myaws"
	"cartoon-gin/utils"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)



var maxIdFile = "../../max_id_to_oss.log"

func main() {
	//limit := flag.Int("limit",1000,"limit")
	begin := flag.Int("begin",0,"begin")
	flag.Parse()
	var lastMaxId int
	fmt.Println("begin:",*begin)
	if *begin >= 0 {
		lastMaxId = *begin
	}else {
		lastMaxId = getMaxIdToOss()
	}
	fmt.Println("max_id:",lastMaxId)

	timeBegin := time.Now()

	//1. 创建OSSClient实例。
	client, err := oss.New(configs.OssEndPoint, configs.OssAccessKeyId, configs.OsaAccessSecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	//1. 获取存储空间。
	bucket, err := client.Bucket(configs.OssBucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 指定存储类型为归档存储。
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	// 指定访问权限为公共读。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	//imgList := dao.FindCartoonsHoverImageForUpload(*limit,lastMaxId)
	imgList := dao.FindImagesForUpload(1000,lastMaxId)
	fmt.Printf("count:%d \n",len(imgList))

	maxId := 0
	for _, row := range imgList {
		t1 := time.Now()
		src := row.ImageAddr
		filename, filebody := myaws.GetFileBodyAndName(src)

		// 上传字符串。
		err = bucket.PutObject(filename, filebody, storageType, objectAcl)
		if err != nil {
			fmt.Printf("error row_id=%v msg=%v \n",row.ID,err.Error())
			continue
		}

		escaped := time.Since(t1)
		fmt.Printf("Successfully uploaded %q to %q -- escaped: %v\n", filename, myaws.S3_BUCKET,escaped)
		if row.ID > maxId {
			maxId = row.ID
			saveMaxIdToOss(maxId)
		}
	}

	totalTime := time.Since(timeBegin)
	fmt.Printf("Done total_time:%v \n",totalTime)
}

func getMaxIdToOss() int {
	c,err := ioutil.ReadFile(maxIdFile)
	utils.CheckError(err)
	i,err := strconv.Atoi(string(c))
	utils.CheckError(err)
	return i
}
func saveMaxIdToOss(id int) {
	b := []byte(strconv.Itoa(id))
	err := ioutil.WriteFile(maxIdFile,b,0644)
	utils.CheckError(err)
}
