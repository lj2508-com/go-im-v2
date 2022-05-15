package ctrl

import (
	"go-im-v2/args"
	"go-im-v2/model"
	"go-im-v2/serivce"
	"go-im-v2/util"
	"net/http"
)

var contactService serivce.ContactService

//添加好友
func Addfriend(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	util.Bind(req, &arg)
	//调用service
	err := contactService.AddFriend(arg.Userid, arg.Dstid)
	//
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, nil, "好友添加成功")
	}
}

//获取好友列表
func LoadFriend(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)

	users := contactService.SearchFriend(arg.Userid)
	util.RespOkList(w, users, len(users))
}

//获取群列表
func LoadCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)
	comunitys := contactService.SearchComunity(arg.Userid)
	util.RespOkList(w, comunitys, len(comunitys))
}

func CreateCommunity(w http.ResponseWriter, req *http.Request) {
	var arg model.Community
	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)
	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, com, "")
	}
}

//新建群
func JoinCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg

	//如果这个用的上,那么可以直接
	util.Bind(req, &arg)
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)
	//加入群之后刷新数据
	AddGroupId(arg.Userid, arg.Dstid)

	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, nil, "")
	}
}
