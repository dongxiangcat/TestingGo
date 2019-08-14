package model

/**
 * xorm连接数据库操作
 */
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {
	Engine, _ = xorm.NewEngine("mysql", "root:root@/fun_api?charset=utf8")
}

func CreateOrmModel() OrmModel {
	return OrmModel{Engine}
}

type OrmModel struct {
	*xorm.Engine
}
