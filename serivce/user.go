package serivce

import (
	"errors"
	"fmt"
	"go-im-v2/model"
	"go-im-v2/util"
	"math/rand"
	"time"
)

type UserService struct {
}

//用户注册函数
func (s *UserService) Register(mobile, plainPass, nickName, avatar, sex string) (user model.User, err error) {
	//先判断用户是否注册过 手机号唯一
	tmp := model.User{}
	_, err = DbEngin.Where("mobile=?", mobile).Get(&tmp)
	if err != nil {
		return tmp, err
	}
	if tmp.Id > 0 {
		return tmp, errors.New("注册失败！账号已存在！")
	}
	//账号不存在则注册账号
	tmp.Mobile = mobile
	tmp.Nickname = nickName
	tmp.Avatar = avatar
	tmp.Sex = sex
	tmp.Createat = time.Now()
	tmp.Salt = fmt.Sprintf("%06d", rand.Int31())
	tmp.Passwd = util.MakePasswd(tmp.Salt, plainPass)

	fmt.Println(tmp)
	DbEngin.InsertOne(tmp)
	//保存成功，返回用户
	return tmp, nil
}

//用户登陆
func (s *UserService) Login(mobile, plainPass, nickName, avatar, sex string) (user model.User, err error) {
	return user, nil
}
