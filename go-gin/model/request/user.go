package request

type LoginUser struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	ValidCode string `json:"validcode" form:"validcode"`
}

type RegisterUser struct {
	Username  string `json:"username" form:"username"`
	Password1 string `json:"password1" form:"password1"`
	Password2 string `json:"password2" form:"password2"`
	ValidCode string `json:"validcode" form:"validcode"`
}
