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

	apiData := entities.Api{Name: name, Desc: desc, Hosts: hosts, Pathinfo: pathinfo}
	_, err := models.ApiOrm.Insert(&apiData)

	var jsonString string
	if err == nil {
		jsonString = utils.EncodeJson(vo.Succ())

	} else {
		jsonString = utils.EncodeJson(vo.Error(1, "插入失败"))
	}
	this.Echo(jsonString)
}

func (this *ApiController) List() {
	draw := this.GetInt("draw")
	start := this.GetInt("start")
	length := this.GetInt("length")

	list := make([]entities.Api, 0)
	models.ApiOrm.Limit(length, start).Find(&list)

	api := new(entities.Api)
	total, _ := models.ApiOrm.Count(api)

	// DataTable 返回格式

	this.Echo(utils.EncodeJson(vo.DataTable(draw, list, int(total))))
}
