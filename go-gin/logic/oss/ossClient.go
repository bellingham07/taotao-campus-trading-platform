package ossLogic

import (
	"com.xpwk/go-gin/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

var OSSClient OssClient
var bucketName string

type OssClient struct {
	oss.Client
}

func InitOSS(config config.OSSConfig) {
	OSSClient, err := oss.New(config.EndPoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		panic(fmt.Sprintf("连接OSS服务失败：%s\n", err.Error()))
	}
	bucketName = config.BucketName
}

func (oc *OssClient) Upload(file multipart.File, objectName string) error {
	// 获取存储空间。
	bucket, err := oc.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return err
	}
	return nil
}

func (oc *OssClient) Delete(objectName string) error {
	// 获取存储空间。
	bucket, err := oc.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 删除文件。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		return err
	}
	return nil
}
