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
	http.ListenAndServe(":8000", nil)
}
