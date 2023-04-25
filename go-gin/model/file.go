package model

import "time"

type FileAsset struct {
	Id        int64
	UserId    int64
	ArticleId int64
	Name      string
	Url       string
	CreateAt  time.Time
}
