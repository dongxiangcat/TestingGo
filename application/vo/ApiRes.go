package vo

/* 输出结构体 */
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Succ() Result {
	return Result{0, "succ"}
}

func Error(code int, msg string) Result {
	return Result{code, msg}
}
