package router

import "net/http"
import "fmt"

// Handler 接口实现
type Router struct {
	Method string
	Path   string
	Hander func(w http.ResponseWriter, r *http.Request)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ri执行了ServeHTTP")
	if router.Method == r.Method {
		router.Hander(w, r)
	} else {
		fmt.Fprintln(w, "Method错误")
	}
}

func NewRouter(method string, path string, hander func(w http.ResponseWriter, r *http.Request)) *Router {
	return &Router{method, path, hander}
}

func (r *Router) Register() {
	http.Handle(r.Path, r)
}
