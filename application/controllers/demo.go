package controllers

import (
	"fmt"
	"test/application/models"
	"test/application/utils"
	"test/core/controller"
)

var Demo *DemoController

type DemoController struct {
	controller.Controller
}

func init() {
	Demo = &DemoController{}
}

func (this *DemoController) Test2() {
	fmt.Println("Test2执行")
}

func (this *DemoController) Test() {

	fmt.Println(controller.ContextInput.GetParam["aa"])
	fmt.Println(controller.ContextInput.PostParam["a"])

	rows, err := models.DemoModel.DB.Query("select name,hosts from api where id = ? ", 1)
	defer rows.Close()

	for rows.Next() {
		var name string
		var hosts string
		err := rows.Scan(&name, &hosts)
		if err != nil {
			fmt.Println(err)
		}
		jsonString := utils.EncodeJson(DemoController{})
		fmt.Println(jsonString)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

}
