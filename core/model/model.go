package model

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	DB    *sql.DB
	Table string
}

func CreateModel(Table string) Model {

	return Model{createConnect(), Table}
}

// 单例connect
var DB *sql.DB

func createConnect() *sql.DB {
	if DB == nil {
		fmt.Println("进入connect初始化")
		db, err := sql.Open("mysql", "root:root@/fun_api?charset=utf8")
		if err != nil {
			fmt.Println("open mysql failed,", err)
			return nil
		}
		DB = db
	}
	return DB
}
