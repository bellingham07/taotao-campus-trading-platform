package ossLogic

import (
	"com.xpwk/go-gin/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"math/rand"
	"mime/multipart"
	"path"
	"strconv"
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

func InitOSS(config config.OSSConfig) {
	endPoint = config.EndPoint
	accessKeyId = config.AccessKeyId
	accessKeySecret = config.AccessKeySecret
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
	err = bucket.PutObject(bucketName, file)
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
	//return nil
	return nil
}

func (oc *OssClient) MultiDelete(objectNames []string) error {
	bucket, err := oc.Bucket(bucketName)
	if err != nil {
		return err
	}
	// 删除单个文件。
	_, err = bucket.DeleteObjects(objectNames, oss.DeleteObjectsQuiet(true))
	if err != nil {
		log.Println("OSS批量删除错误：" + err.Error())
		return err
	}
	return nil
}

func generateURL(filename string, belong string) string {
	suffix := path.Ext(filename)
	filename = strings.TrimSuffix(filename, suffix)
	randomInt := rand.Int()
	return "images/" + belong + "/" + filename + strconv.Itoa(randomInt) + time.Now().String() + suffix
}
