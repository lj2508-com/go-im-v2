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
	//根据实体来建立表
	DbEngin.Sync2(new(model.User), new(model.Contact), new(model.Community))
	DbEngin.SetMaxOpenConns(200)
	fmt.Println("== 数据库初始化成功 ==")
}
