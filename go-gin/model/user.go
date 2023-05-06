package model

import "time"

type UserInfo struct {
	Id        int64     `json:"id" form:"id"`             // bigint，id(snowflake)
	Username  string    `json:"username" form:"username"` // 账户
	Password  string    `json:"password" form:"password"` // 密码
	Name      string    `json:"name" form:"name"`         // 姓名
	Gender    string    `json:"gender" form:"gender"`     // 性别
	Phone     string    `json:"phone" form:"phone"`       // 手机号
	Avatar    string    `json:"avatar"`                   // 头像url
	Intro     string    `json:"intro"`                    // 个人简介
	Location  string    `json:"location"`                 // 住址
	LastLogin time.Time `json:"lastLogin"`                // date，上次登录时间
	Status    int64     `json:"status"`                   // 用户账户状态
	Call      string    `json:"call"`                     // 称号
	Done      int64     `json:"done"`                     // 成交数
	Fans      int64     `json:"fans"`                     // 粉丝数
	Follow    int64     `json:"follow"`                   // 关注数
	Like      int64     `json:"like"`                     // 点赞数
	Negative  int64     `json:"negative"`                 // 差评次数
	Positive  int64     `json:"positive"`                 // 好评次数
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

type UserLocation struct {
	Id       int
	Name     string
	createBy int64
	createAt time.Time
	updateBy int64
	updateAt time.Time
}
