package ossLogic

import (
	"com.xpdj/go-gin/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

var (
	OSSClient       = new(OssClient)
	bucketName      string
	endPoint        string
	accessKeyId     string
	accessKeySecret string
	baseUrl         string
)

type OssClient struct {
	*oss.Client
}

func InitOSS(config *config.OSSConfig) {
	endPoint = config.EndPoint
	accessKeyId = config.AccessKeyId
	accessKeySecret = config.AccessKeySecret
	bucketName = config.BucketName
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(fmt.Sprintf("连接OSS服务失败：%s\n", err.Error()))
	}
	OSSClient.Client = client
	baseUrl = "https://" + bucketName + "." + endPoint + "/"
}

func (oc *OssClient) Upload(fileHeader *multipart.FileHeader, userId string) (string, string, error) {
	// 获取bucket
	bucket, err := oc.Bucket(bucketName)
	if err != nil {
		return "", "", err
	}
	// 生成url
	filename := fileHeader.Filename
	url := generateURL(filename, userId)
	// 上传文件。
	file, err := fileHeader.Open()
	if err != nil {
		return "", "", err
	}
	err = bucket.PutObject(url, file)
	if err != nil {
		return "", "", err
	}
	return baseUrl + url, url, nil
}

func (oc *OssClient) MultiUpload(fileHeaders []*multipart.FileHeader, userId string) ([]string, []string, error) {
	var urls, objectNames []string
	for _, fileHeader := range fileHeaders {
		url, objectName, err := oc.Upload(fileHeader, userId)
		if err != nil {
			return nil, nil, err
		}
		urls = append(urls, url)
		objectNames = append(objectNames, objectName)
	}
	return urls, objectNames, nil
}

func (oc *OssClient) Delete(objectName string) {
	// 获取存储空间。
	bucket, _ := oc.Bucket(bucketName)
	// 删除文件。
	// objectName表示删除OSS文件时需要指定包含文件后缀，不包含Bucket名称在内的完整路径，例如exampledir/exampleobject.txt。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	if err := bucket.DeleteObject(objectName); err != nil {
		log.Println("Delete 文件删除失败（OSS）" + err.Error())
	}
}

func (oc *OssClient) MultiDelete(objectNames []string) {
	bucket, _ := oc.Bucket(bucketName)
	if _, err := bucket.DeleteObjects(objectNames, oss.DeleteObjectsQuiet(true)); err != nil {
		log.Println("MultiDelete 文件批量删除失败（OSS）" + err.Error())
	}
}

func generateURL(filename string, belong string) string {
	suffix := path.Ext(filename)
	filename = strings.TrimSuffix(filename, suffix)
	return "images/" + belong + "/" + time.Now().String() + "/" + filename + time.Now().String() + suffix
}
