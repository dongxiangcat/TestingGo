package router

import (
	"net/http"
	"test/core/controller"
)
import "fmt"

import "reflect"

import "net/url"

// Handler 接口实现

type Router struct {
	Path string
	Cr   *ControllerRegister
}

type ControllerRegister struct {
	Method string
	Hander interface{}
}

func (cr *ControllerRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("router执行了ServeHTTP")
	if cr.Method == r.Method || cr.Method == "*" {

		// 请求参数
		controller.ContextInput = controller.NewInput()
		controller.ContextInput.GetParam = map[string]interface{}{}
		controller.ContextInput.PostParam = map[string]interface{}{}

		// GET 参数赋值
		u, _ := url.Parse(r.URL.String())
		getForm, _ := url.ParseQuery(u.RawQuery)

		for k, param := range getForm {
			controller.ContextInput.GetParam[k] = param
		}

		// POST参数赋值
		r.ParseForm()
		postForm := r.PostForm
		for k, param := range postForm {
			controller.ContextInput.PostParam[k] = param
		}
		// 调用方法
		f := reflect.ValueOf(cr.Hander)
		in := make([]reflect.Value, 0)
		f.Call(in)
	} else {
		fmt.Fprintln(w, "Method错误")
	}
}

func NewControllerRegister(method string, hander interface{}) *ControllerRegister {
	return &ControllerRegister{method, hander}
}

func (r *Router) RegisterWithMethod(path string, hander interface{}, method string) {
	r.Path = path
	r.Cr = NewControllerRegister(method, hander)
	http.Handle(r.Path, r.Cr)
}

func (r *Router) Register(path string, hander interface{}) {
	r.Path = path
	r.Cr = NewControllerRegister("*", hander)
	http.Handle(r.Path, r.Cr)
}
