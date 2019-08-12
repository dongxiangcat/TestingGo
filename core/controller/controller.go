package controller

type Controller struct {
	Data           map[interface{}]interface{}
	controllerName string
	methodName     string
}

func InitController() Controller {
	var controller Controller
	return controller
}
