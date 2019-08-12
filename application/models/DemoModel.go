package models

import (
	"test/core/model"
)

type demoModel struct {
	model.Model
}

// demoModel指针
var DemoModel *demoModel

func init() {
	DemoModel = &demoModel{model.CreateModel("api")}
}
