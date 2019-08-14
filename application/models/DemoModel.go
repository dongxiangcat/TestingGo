package models

import (
	"test/core/model"
)

type demoModel struct {
	model.DBModel
}

// demoModel指针
var DemoModel *demoModel

func init() {
	DemoModel = &demoModel{model.CreateModel("api")}
}
