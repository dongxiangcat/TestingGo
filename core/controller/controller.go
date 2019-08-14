package controller

import (
	"fmt"
	"net/http"
)

// controller提供一些快捷便利的输出与模板服务
type Controller struct {
}

func (c *Controller) GetString(key string) string {
	var temp interface{} = ContextInput.GetParam[key]
	if temp == nil {
		return ""
	}
	return temp.(string)
}

func (c *Controller) PostString(key string) string {
	var temp interface{} = ContextInput.PostParam[key]
	if temp == nil {
		return ""
	}
	return temp.(string)
}

func (c *Controller) Echo(param interface{}) {
	fmt.Fprintf(ContextOutPut.w, "%s", param)
}

// Input提供web数据的绑定
var ContextInput *Input

type Input struct {
	GetParam  map[string]interface{}
	PostParam map[string]interface{}
}

func NewInput() *Input {
	ContextInput = new(Input)
	ContextInput.GetParam = map[string]interface{}{}
	ContextInput.PostParam = map[string]interface{}{}
	return ContextInput
}

/*  输出  */
var ContextOutPut *OutPut

type OutPut struct {
	w http.ResponseWriter
}

func NewOutPut(w http.ResponseWriter) *OutPut {
	ContextOutPut = new(OutPut)
	ContextOutPut.w = w
	return ContextOutPut
}
