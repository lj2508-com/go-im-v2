package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", userlogin)
	http.ListenAndServe(":8090", nil)
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
