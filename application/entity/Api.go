package entities

import (
	"time"
)

type Api struct {
	Id         int64
	Name       string    `xorm:"varchar(25) notnull"`
	Desc       string    `xorm:"varchar(255) notnull"`
	Hosts      string    `xorm:"varchar(255) notnull"`
	Pathinfo   string    `xorm:"varchar(255) notnull"`
	Method     string    `xorm:"varchar(10) notnull"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}
