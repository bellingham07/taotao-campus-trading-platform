package response

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

const (
	SUCCESS = "success"
	ERROR   = "error"
)
