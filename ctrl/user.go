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
	user, err := userService.Login(mobile, passwd)
	if err != nil {
		util.RespFail(writer, err.Error())
	} else {
		util.RespOk(writer, user, "")
	}
}
