package controller

// 指针
var BaseController *Controller

type Controller struct {
	Data           map[interface{}]interface{}
	GetData        map[interface{}]interface{}
	PostData       map[interface{}]interface{}
	controllerName string
	methodName     string
}

func GetControllerPoint() *Controller {
	if BaseController == nil {
		BaseController = new(Controller)
		BaseController.Init()
	}
	return BaseController
}

func (c *Controller) Init() {
	c.Data = map[interface{}]interface{}{}
	c.GetData = map[interface{}]interface{}{}
	c.PostData = map[interface{}]interface{}{}
}
