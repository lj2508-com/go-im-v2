package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"xorm.io/xorm"
)

var DbEngin *xorm.Engine

func main() {
	http.HandleFunc("/user/login", userlogin)
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	registerView()
	http.ListenAndServe(":8090", nil)
}

func Ahuiafhia() {

}

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
func registerView() {
	glob, err := template.ParseGlob("view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range glob.Templates() {
		name := v.Name()
		http.HandleFunc(name, func(writer http.ResponseWriter, request *http.Request) {
			glob.ExecuteTemplate(writer, name, nil)
		})
	}
}
func userlogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
	loginSu := false
	if mobile == "18100000000" && passwd == "123456" {
		loginSu = true
	}
	if loginSu {
		data := make(map[string]interface{})
		data["id"] = 1
		data["token"] = "test"
		resp(writer, 0, data, "")
	} else {
		resp(writer, -1, nil, "账号或密码错误！")
	}
}

type respBady struct {
	Code int
	Data interface{}
	Msg  string
}

func resp(writer http.ResponseWriter, code int, data interface{}, msg string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	bady := respBady{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	json, _ := json.Marshal(bady)
	writer.Write(json)
}