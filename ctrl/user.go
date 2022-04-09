package ctrl

import (
	"go-im-v2/serivce"
	"go-im-v2/util"
	"net/http"
)

var userService serivce.UserService

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	plainPass := request.PostForm.Get("passwd")
	sex := request.PostForm.Get("sex")
	nickName := request.PostForm.Get("nickName")
	avatar := request.PostForm.Get("avatar")

	user, err := userService.Register(mobile, plainPass, nickName, avatar, sex)
	if err == nil {
		util.RespOk(writer, user, "账号注册成功！")
	} else {
		util.RespFail(writer, err.Error())
	}
}

func UserLogin(writer http.ResponseWriter, request *http.Request) {
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
		util.RespOk(writer, data, "")
	} else {
		util.RespFail(writer, "账号或密码错误！")
	}
}
