package entities

import (
	"time"
)

type ApiManager struct {
	Id         int64
	Name       string    `xorm:"varchar(25) notnull"`
	Desc       string    `xorm:"varchar(255) notnull"`
	ApiList    string    `xorm:"varchar(255) notnull"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
