package router

import (
	"net/http"
	"reflect"
	"test/core/controller"
)
import "fmt"

import "net/url"

// 注册器
type Router struct {
	Path string
	Cr   *ControllerRegister
}

// Handler 接口实现 ControllerRegister
type ControllerRegister struct {
	Method string
	Hander interface{}
}

func (cr *ControllerRegister) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("router执行了ServeHTTP")
	if cr.Method == r.Method || cr.Method == "*" {

		var paramString interface{}

		// 请求参数
		controller.ContextInput = controller.NewInput()
		controller.ContextOutPut = controller.NewOutPut(w)

		// GET 参数赋值
		u, _ := url.Parse(r.URL.String())
		getForm, _ := url.ParseQuery(u.RawQuery)

		for k, param := range getForm {
			// 必须先定义一个interface{}变量,才能类型转换
			paramString = param
			controller.ContextInput.GetParam[k] = paramString.([]string)[0]
		}

		// POST参数赋值
		r.ParseForm()
		postForm := r.PostForm

		for k, param := range postForm {
			// 必须先定义一个interface{}变量,才能类型转换
			paramString = param
			controller.ContextInput.PostParam[k] = paramString.([]string)[0]
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

/* 路由注册 携带GET/POST限制 */
func (r *Router) RegisterWithMethod(path string, hander interface{}, method string) {
	r.Path = path
	r.Cr = NewControllerRegister(method, hander)
	http.Handle(r.Path, r.Cr)
}

/* 路由注册 */
func (r *Router) Register(path string, hander interface{}) {
	r.Path = path
	r.Cr = NewControllerRegister("*", hander)
	http.Handle(r.Path, r.Cr)
}
