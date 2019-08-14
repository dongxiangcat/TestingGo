package controllers

import (
	"test/application/entity"
	"test/application/models"
	"test/application/utils"
	"test/application/vo"
	"test/core/controller"
)

/*
 * 初始化三件套START
 */
var Api *ApiController

type ApiController struct {
	controller.Controller
}

func init() {
	Api = &ApiController{}
}

/*
 * 初始化三件套END
 */

func (this *ApiController) Create() {

	name := this.PostString("name")
	desc := this.PostString("desc")
	hosts := this.PostString("hosts")
	pathinfo := this.PostString("pathinfo")

	apiModel := entities.Api{Name: name, Desc: desc, Hosts: hosts, Pathinfo: pathinfo}
	_, err := models.ApiOrm.Insert(&apiModel)

	var jsonString string
	if err == nil {
		jsonString = utils.EncodeJson(vo.Succ())

	} else {
		jsonString = utils.EncodeJson(vo.Error(1, "插入失败"))
	}
	this.Echo(jsonString)
}
