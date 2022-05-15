package serivce

import (
	"errors"
	"go-im-v2/model"
	"time"
)

type ContactService struct {
}

//添加好友
func (service *ContactService) AddFriend(userid, dstid int64) error {
	if userid == dstid {
		return errors.New("不能添加自己为好友！")
	}
	//判断是否已经加了好友
	tmp := model.Contact{}
	DbEngin.Where("ownerid = ?", userid).And("dstid=?", dstid).And("cate = ?", model.CONCAT_CATE_USER).Get(&tmp)
	if tmp.Id != 0 {
		return errors.New("你已添加对方为好友，不能重复添加！")
	}
	//新增好友

	//开启事务。防止部分成功
	session := DbEngin.NewSession()
	session.Begin()
	_, err1 := session.InsertOne(model.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//插对方的
	_, err2 := session.InsertOne(model.Contact{
		Ownerid:  dstid,
		Dstobj:   userid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})

	if err1 == nil && err2 == nil {
		session.Commit()
		return nil
	} else {
		session.Rollback()
		if err1 != nil {
			return err1
		} else {
			return err2
		}
	}
}

//查找好友
func (service *ContactService) SearchFriend(userId int64) []model.User {
	conconts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DbEngin.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&conconts)
	for _, v := range conconts {
		objIds = append(objIds, v.Dstobj)
	}
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	DbEngin.In("id", objIds).Find(&coms)
	return coms
}

func (service *ContactService) SearchComunityIds(userId int64) (comIds []int64) {
	//todo 获取用户全部群ID
	conconts := make([]model.Contact, 0)
	comIds = make([]int64, 0)

	DbEngin.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	return comIds
}

//搜索群
func (service *ContactService) SearchComunity(userId int64) []model.Community {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	DbEngin.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.Dstobj)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	DbEngin.In("id", comIds).Find(&coms)
	return coms
}

//加群
func (service *ContactService) JoinCommunity(userId, comId int64) error {
	cot := model.Contact{
		Ownerid: userId,
		Dstobj:  comId,
		Cate:    model.CONCAT_CATE_COMUNITY,
	}
	DbEngin.Get(&cot)
	if cot.Id == 0 {
		cot.Createat = time.Now()
		_, err := DbEngin.InsertOne(cot)
		return err
	} else {
		return nil
	}

}

//建群
func (service *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.Ownerid == 0 {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	num, err := DbEngin.Count(&com)

	if num > 5 {
		err = errors.New("一个用户最多只能创见5个群")
		return com, err
	} else {
		comm.Createat = time.Now()
		session := DbEngin.NewSession()
		session.Begin()
		_, err = session.InsertOne(&comm)
		if err != nil {
			session.Rollback()
			return com, err
		}
		_, err = session.InsertOne(
			model.Contact{
				Ownerid:  comm.Ownerid,
				Dstobj:   comm.Id,
				Cate:     model.CONCAT_CATE_COMUNITY,
				Createat: time.Now(),
			})
		if err != nil {
			session.Rollback()
		} else {
			session.Commit()
		}
		return com, err
	}
}
