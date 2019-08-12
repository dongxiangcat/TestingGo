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
	Demo = &DemoController{*(controller.GetControllerPoint())}
}

func (this *DemoController) Test2() {
	fmt.Println("Test2执行")
}

func (this *DemoController) Test() {

	rows, err := models.DemoModel.DB.Query("select name,hosts from api where id = ? ", 1)
	defer rows.Close()

	for rows.Next() {
		var name string
		var hosts string
		err := rows.Scan(&name, &hosts)
		if err != nil {
			fmt.Println(err)
		}
		jsonString, _ := utils.EncodeJson(DemoController{})
		fmt.Println(jsonString)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

}
