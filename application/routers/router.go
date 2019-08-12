package routers

import (
	"test/application/controllers"
	"test/core/router"
)

func init() {
	var routerPoint = &router.Router{}
	routerPoint.Register("/", controllers.Demo.Test)
}
