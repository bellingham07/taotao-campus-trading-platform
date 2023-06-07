// Code generated by goctl. DO NOT EDIT.
package types

type IdReq struct {
	Id int64 `json:"id"`
}

type BaseResp struct {
	Code int8   `json:"code"`
	Msg  string `json:"msg"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username  string `json:"username"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

type UserInfoReq struct {
	IdReq
	Username string `json:"username"` // 账户
	Name     string `json:"name"`     // 姓名
	Gender   string `json:"gender"`   // 性别
	Phone    string `json:"phone"`    // 手机号
	Avatar   string `json:"avatar"`   // 头像url
	Intro    string `json:"intro"`    // 个人简介
	Location string `json:"location"` // 住址
}

type LoginResp struct {
	Token string `json:"token"`
}

type RegisterResp struct {
	BaseResp
}

type UserInfoResp struct {
	IdReq
	Username  string `json:"username"`  // 账户
	Name      string `json:"name"`      // 姓名
	Gender    string `json:"gender"`    // 性别
	Phone     string `json:"phone"`     // 手机号
	Avatar    string `json:"avatar"`    // 头像url
	Intro     string `json:"intro"`     // 个人简介
	Location  string `json:"location"`  // 住址
	LastLogin string `json:"lastLogin"` // date，上次登录时间
	Status    int64  `json:"status"`    // 用户账户状态
	Call      string `json:"call"`      // 称号
	Done      int64  `json:"done"`      // 成交数
	Fans      int64  `json:"fans"`      // 粉丝数
	Follow    int64  `json:"follow"`    // 关注数
	Like      int64  `json:"like"`      // 点赞数
	Negative  int64  `json:"negative"`  // 差评次数
	Positive  int64  `json:"positive"`  // 好评次数
}

type DormResp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
