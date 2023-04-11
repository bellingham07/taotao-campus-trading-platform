package model

import "time"

type User struct {
	Id        int64
	Username  string
	Password  string
	Gender    string
	Phone     string
	Avatar    string
	Intro     string
	LastLogin time.Time
	Like      int32
	Status    int8
	Done      int32
	Call      string
	Fans      int32
	Follow    int32
	Positive  int32
	Negative  int32
}

type UserCall struct {
	Id       int32
	Name     string
	CreateBy int64
	CreateAt time.Time
	UpdateBy int64
	UpdateAt time.Time
}

type UserCollect struct {
	Id            int64
	UserId        int64
	CollectUserId int64
	CreateAt      time.Time
}