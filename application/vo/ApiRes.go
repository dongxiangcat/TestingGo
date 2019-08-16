package vo

/* 输出结构体 */
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Succ() Result {
	return Result{Code: 0, Message: "succ"}
}

func Data(data interface{}) Result {
	return Result{0, "succ", data}
}

func Error(code int, msg string) Result {
	return Result{Code: code, Message: msg}
}
