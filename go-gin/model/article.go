package model

import "time"

type ArticlePost struct {
	Id        int64
	UserId    int64
	Title     string
	Content   string
	Type      int8
	CommentId int64
	Status    int8
	CreateAt  time.Time
	UpdateAt  time.Time
}

type ArticleComment struct {
	Id       int64
	UserId   int64
	Content  string
	RootId   int64
	ToUserId int64
	CreateAt time.Time
}
