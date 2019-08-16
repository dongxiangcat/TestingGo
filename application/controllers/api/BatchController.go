package apis

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"test/application/entity"
	"test/application/models"
	appUtils "test/application/utils"
	"test/core/controller"
	"test/core/utils"
)

/*
 * 初始化三件套START
 */
var Batch *BatchController

type BatchController struct {
	controller.Controller
}

func init() {
	Batch = &BatchController{}
}

/*
 * 初始化三件套END
 */
var wg sync.WaitGroup

// 协程请求多个API
func (this *BatchController) HttpRequestBatch() {
	// 接口id
	ids := this.GetString("ids")
	arr := strings.Split(ids, ",")

	var apiList []entities.Api = make([]entities.Api, 0)
	models.ApiOrm.In("id", arr).Find(&apiList)

	wg.Add(len(apiList))
	for _, api := range apiList {
		// 获取参数 JSON
		param := this.PostString(strconv.Itoa(int(api.Id))) // 数字变成字符串
		// 转为map类型
		paramMap := make(map[string]interface{}, 0)
		appUtils.DecodeJson([]byte(param), &paramMap)

		go this.HttpRequestSync(api.Hosts+api.Pathinfo, api.Method, paramMap)
	}

	wg.Wait()
}

func (this *BatchController) HttpRequestSync(url string, method string, param map[string]interface{}) {
	var result string
	if method == "GET" {
		result = utils.Get(url, param)
	} else {
		result = utils.Post(url, param)
	}
	fmt.Println("结果:" + result)
	this.Echo(result)
	wg.Done()
}

/** 同步请求多接口 做对比用 **/
func (this *BatchController) HttpRequestLine() {
	ids := this.GetString("ids")
	arr := strings.Split(ids, ",")

	var apiList []entities.Api = make([]entities.Api, 0)
	models.ApiOrm.In("id", arr).Find(&apiList)

	for _, api := range apiList {
		// 获取参数 JSON
		param := this.PostString(strconv.Itoa(int(api.Id))) // 数字变成字符串
		// 转为map类型
		paramMap := make(map[string]interface{}, 0)
		appUtils.DecodeJson([]byte(param), &paramMap)
		this.HttpRequest(api.Hosts+api.Pathinfo, api.Method, paramMap)
	}
}

func (this *BatchController) HttpRequest(url string, method string, param map[string]interface{}) {
	var result string
	if method == "GET" {
		result = utils.Get(url, param)
	} else {
		result = utils.Post(url, param)
	}
	fmt.Println("结果:" + result)
	this.Echo(result)
}
