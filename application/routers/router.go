package routers

import (
	"test/application/controllers"
	"test/core/router"
)

func init() {
	router.NewRouter("GET", "/", controllers.Demo.Test).Register()
}
