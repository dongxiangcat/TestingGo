package controller

// 指针
var BaseController *Controller

// controller提供一些快捷便利的输出与模板服务
type Controller struct {
}

// Input提供web数据的绑定
var ContextInput *Input

type Input struct {
	GetParam  map[string]interface{}
	PostParam map[string]interface{}
}

func NewInput() *Input {
	return new(Input)
}
