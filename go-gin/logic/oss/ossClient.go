package ossLogic

import (
	"com.xpwk/go-gin/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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
	return "https://" + bucketName + "." + endPoint + url, filename, nil
}

func (oc *OssClient) MultiUpload(fileHeaders []*multipart.FileHeader, userId string) ([]string, []string, error) {
	var urls, filenames []string
	for _, fileHeader := range fileHeaders {
		url, filename, err := oc.Upload(fileHeader, userId)
		if err != nil {
			return nil, nil, err
		}
		urls = append(urls, url)
		filenames = append(filenames, filename)
	}
	return urls, filenames, nil
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

func generateURL(file string, belong string) string {
	suffix := path.Ext(file)
	filename := strings.TrimSuffix(file, suffix)
	randomInt := rand.Int()
	url := "/images/" + belong + "/" + time.Now().String() + "/" + filename + strconv.Itoa(randomInt) + suffix
	return url
}
