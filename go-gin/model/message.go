package model

import "time"

type MessageTemplate struct {
	Id   int
	Name string
}

type MessageContent struct {
	Id       int
	ToUserId int64
	Content  string
	time     time.Time
}

type MessageChat struct {
	Id      int64
	Content string
	time.Time
}
