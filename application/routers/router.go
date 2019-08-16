package routers

import (
	"test/application/controllers"
	"test/application/controllers/api"
	"test/core/router"
)

func init() {
	var routerPoint = &router.Router{}
	routerPoint.Register("/api/create", controllers.Api.Create)
	routerPoint.Register("/api/list", controllers.Api.List)
	routerPoint.Register("/out/api/batch/", apis.Batch.HttpRequestBatch)
	routerPoint.Register("/out/api/line/", apis.Batch.HttpRequestLine)

	routerPoint.Register("/api/test/", controllers.Demo.Test2)
}
