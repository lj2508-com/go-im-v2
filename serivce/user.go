package serivce

import (
	"go-im-v2/model"
)

type UserService struct {
}

//用户注册函数
func (s *UserService) Register(mobile, plainPass, nickName, avatar, sex string) (user model.User, err error) {
	//先判断用户是否注册过 手机号唯一

	return user, nil
}

//用户登陆
func (s *UserService) Login(mobile, plainPass, nickName, avatar, sex string) (user model.User, err error) {
	return user, nil
}
