package model

import "time"

type (
	UserInfo struct {
		Id       int64  `json:"id"  xorm:"id"`             // id(snowflake)
		Username string `json:"username"  xorm:"username"` // 账户
		Password string `json:"password" xorm:"password"`  // 密码
		Name     string `json:"name" xorm:"name"`          // 姓名x
		Gender   int64  `json:"gender" xorm:"gender"`      // 性别
		Phone    string `json:"phone" xorm:"phone"`        // 手机号
		Avatar   string `json:"avatar" xorm:"avatar"`      // 头像url
		Intro    string `json:"intro" xorm:"intro"`        // 个人简介
		Location string `json:"location" xorm:"location"`  // 住址
		Like     int64  `json:"like" xorm:"like"`          // 获赞数
		Status   int64  `json:"status" xorm:"status"`      // 用户账户状态
		Done     int64  `json:"done" xorm:"done"`          // 成交数
		Call     string `json:"call" xorm:"call"`          // 称号
		Fans     int64  `json:"fans" xorm:"fans"`          // 粉丝数
		Follow   int64  `json:"follow" xorm:"follow"`      // 关注数
		Positive int64  `json:"positive" xorm:"positive"`  // 好评次数
		Negative int64  `json:"negative" xorm:"negative"`  // 差评次数
	}

	UserFollow struct {
		Id           int64     `json:"id" xorm:"id"`                       // id
		UserId       int64     `json:"userId" xorm:"user_id"`              // 用户id
		FollowUserId int64     `json:"followUserId" xorm:"follow_user_id"` // 收藏的用户id
		CreateAt     time.Time `json:"createAt" xorm:"create_at"`          // 创建时间
	}

	UserLocation struct {
		Id       int       `json:"id" xorm:"id"`              // id
		Name     string    `json:"name" xorm:"name"`          // 地址名
		CreateBy int64     `json:"createBy" xorm:"create_by"` // 管理员的id
		CreateAt time.Time `json:"createAt" xorm:"create_at"` // 创建时间
		UpdateBy int64     `json:"updateBy" xorm:"update_by"` // 管理员的id
		UpdateAt time.Time `json:"updateAt" xorm:"update_at"` // 更新时间
	}

	UserCall struct {
		Id       int64     `json:"id" xorm:"id"`              // id
		Name     string    `json:"name" xorm:"name"`          // 称号名字
		CreateBy int64     `json:"createBy" xorm:"create_by"` // 管理员的id
		CreateAt time.Time `json:"createAt" xorm:"create_at"` // 创建时间
		UpdateBy int64     `json:"updateBy" xorm:"update_by"` // 管理员的id
		UpdateAt time.Time `json:"updateAt" xorm:"update_at"` // 更新时间
	}

	UserOpt struct {
		Id       int64  `json:"id" xorm:"id"`
		Username string `json:"username" xorm:"username"`
		Password string `json:"password" xorm:"password"`
		Name     string `json:"name" xorm:"name"`
		Status   int64  `json:"status" xorm:"status"`
		Level    int64  `json:"level" xorm:"level"`
	}
)
