package ossLogic

import (
	"com.xpwk/go-gin/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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

// TODO
func (oc *OssClient) Upload(fileHeader *multipart.FileHeader, userId string) (string, error) {

	// 获取bucket
	bucket, err := oc.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	// 分离文件名和后缀
	filename, suffix := resolveFilename(fileHeader.Filename)
	url := "/images/" + userId + "/" + filename + time.Now().String() + suffix
	// 上传文件。
	err = bucket.PutObject(bucketName, fileHeader.o)
	if err != nil {
		return "", err
	}
	return "https://" + bucketName + "." + endPoint + url, nil
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
func resolveFilename(file string) (string, string) {
	suffix := path.Ext(file)
	file = strings.TrimSuffix(file, suffix)
	return file, suffix
}
