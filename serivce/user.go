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
	tmp.Passwd = util.MakePasswd(plainPass, tmp.Salt)

	fmt.Println(tmp)
	DbEngin.InsertOne(tmp)
	//保存成功，返回用户
	return tmp, nil
}

//用户登陆
func (s *UserService) Login(mobile, plainPass string) (user model.User, err error) {
	temp := model.User{}
	_, err = DbEngin.Where("mobile=?", mobile).Get(&temp)
	if err != nil {
		return temp, err
	}
	if temp.Id == 0 {
		return temp, errors.New("登陆失败！账号不存在！")
	}
	if !util.ValidatePasswd(plainPass, temp.Salt, temp.Passwd) {
		return temp, errors.New("登陆失败！密码错误！")
	}
	t := fmt.Sprintf("%d", time.Now().Unix())
	temp.Token = util.Md5Encode(t)
	DbEngin.ID(temp.Id).Cols("token").Update(&temp)
	return temp, nil
}

func (s *UserService) Find(userId int64) model.User {
	temp := model.User{}
	DbEngin.ID(userId).Get(&temp)
	return temp
}
