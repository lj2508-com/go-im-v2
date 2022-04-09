package serivce

import (
	"errors"
	"fmt"
	"go-im-v2/model"
	"log"
	"xorm.io/xorm"
)

var DbEngin *xorm.Engine

func init() {
	drivename := "sqlite3"
	dbPath := `/Users/lijiang/Documents/sqllite3.db`
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(drivename, dbPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	DbEngin.ShowSQL(true)
	DbEngin.Sync2(new(model.User))
	DbEngin.SetMaxOpenConns(200)
	fmt.Println("== 数据库初始化成功 ==")
}
