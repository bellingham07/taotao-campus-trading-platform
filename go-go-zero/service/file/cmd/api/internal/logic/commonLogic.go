package logic

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/file/cmd/api/internal/svc"
	"log"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

type CommonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommonLogic {
	return &CommonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommonLogic) Upload(fileHeader *multipart.FileHeader, userId string) (string, string, error) {
	// 获取bucket
	bucket, err := l.svcCtx.Oss.Client.Bucket(l.svcCtx.Oss.BucketName)
	if err != nil {
		return "", "", err
	}
	// 生成url
	filename := fileHeader.Filename
	objectName := genObjectName(filename, userId)
	// 上传文件。
	file, err := fileHeader.Open()
	if err != nil {
		logx.Debugf("[OSS ERROR] Upload 解析图片错误 " + err.Error())
		return "", "", err
	}
	err = bucket.PutObject(objectName, file)
	if err != nil {
		logx.Debugf("[OSS ERROR] Upload 上传图片错误 " + err.Error())
		return "", "", err
	}
	return l.svcCtx.Oss.BaseUrl + objectName, objectName, nil
}

func (l *CommonLogic) MultiUpload(fileHeaders []multipart.FileHeader, userId string) ([]string, []string, error) {
	var urls, objectNames []string
	for _, fileHeader := range fileHeaders {
		url, objectName, err := l.Upload(&fileHeader, userId)
		if err != nil {
			return nil, nil, err
		}
		urls = append(urls, url)
		objectNames = append(objectNames, objectName)
	}
	return urls, objectNames, nil
}

func (l *CommonLogic) Delete(objectName string) {
	// 获取存储空间。
	bucket, _ := l.svcCtx.Oss.Client.Bucket(l.svcCtx.Oss.BucketName)
	// 删除文件。
	// objectName表示删除OSS文件时需要指定包含文件后缀，不包含Bucket名称在内的完整路径，例如exampledir/exampleobject.txt。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	if err := bucket.DeleteObject(objectName); err != nil {
		log.Println("Delete 文件删除失败（OSS）" + err.Error())
	}
}

func (l *CommonLogic) MultiDelete(objectNames []string) {
	bucket, _ := l.svcCtx.Oss.Client.Bucket(l.svcCtx.Oss.BucketName)
	if _, err := bucket.DeleteObjects(objectNames, oss.DeleteObjectsQuiet(true)); err != nil {
		log.Println("MultiDelete 文件批量删除失败（OSS）" + err.Error())
	}
}

func genObjectName(filename string, belong string) string {
	suffix := path.Ext(filename)
	filename = strings.TrimSuffix(filename, suffix)
	t := time.Now()
	fragmt1 := "images/" + t.Format("2006-01") + "/" + t.Format("02")
	fragmt2 := "/" + belong + "/" + filename + time.Now().Format("15:04:05") + suffix
	return fragmt1 + fragmt2
}
