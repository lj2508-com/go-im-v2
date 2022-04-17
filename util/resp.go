package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
	Rows  interface{} `json:"rows,omitempty"`
	Total interface{} `json:"total,omitempty"`
}

func RespOkList(w http.ResponseWriter, lists interface{}, total interface{}) {
	//分页数目,
	RespList(w, 0, lists, total)
}
func RespList(w http.ResponseWriter, code int, data interface{}, total interface{}) {

	w.Header().Set("Content-Type", "application/json")
	//设置200状态
	w.WriteHeader(http.StatusOK)
	//输出
	//定义一个结构体
	//满足某一条件的全部记录数目
	//测试 100
	//20
	h := H{
		Code:  code,
		Rows:  data,
		Total: total,
	}
	//将结构体转化成JSOn字符串
	ret, err := json.Marshal(h)
	if err != nil {
		log.Println(err.Error())
	}
	//输出
	w.Write(ret)
}

func RespFail(w http.ResponseWriter, msg string) {
	resp(w, -1, nil, msg)
}
func RespOk(w http.ResponseWriter, data interface{}, msg string) {
	resp(w, 0, data, msg)
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

type respBady struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
