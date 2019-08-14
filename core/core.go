package core

import (
	"fmt"
	"net/http"
	_ "test/core/router"
)

func init() {
	fmt.Println("引入了core.go")
}

func Run() {
	// 静态资源
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8000", nil)
}
