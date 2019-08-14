package models

import (
	"test/core/model"
)

/* 初始化三件套 */
var ApiOrm *ApiOrmModel

func init() {
	ApiOrm = &ApiOrmModel{model.CreateOrmModel()}
}

type ApiOrmModel struct {
	model.OrmModel
}

/* 初始化三件套END */
