package controllers

import (
	"fmt"
	"net/http"
	"test/application/models"
	"test/application/utils"
	"test/core/controller"
)

var Demo *DemoController

type DemoController struct {
	controller.Controller
}

func init() {
	Demo = &DemoController{controller.InitController()}
}

func (demo *DemoController) Test(w http.ResponseWriter, r *http.Request) {

	fmt.Println((*demo).Data)

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
		fmt.Fprint(w, jsonString, hosts)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

}
