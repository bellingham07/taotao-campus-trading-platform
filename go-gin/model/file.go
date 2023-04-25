package model

import "time"

type FileAsset struct {
	Id         int64
	UserId     int64
	ArticleId  int64
	ObjectName string
	Url        string
	CreateAt   time.Time
}
