package serivce

import (
	"fmt"
	"log"
	"xorm.io/xorm"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "sqlite3"
	dbPath := `/Users/lijiang/Documents/sqllite3.db`
	DbEngin, err := xorm.NewEngine(drivename, dbPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	DbEngin.ShowSQL(true)
	DbEngin.SetMaxOpenConns(200)
	fmt.Println("== 数据库初始化成功 ==")
}
