package request

type LoginUser struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	ValidCode string `json:"validcode" form:"validcode"`
}
