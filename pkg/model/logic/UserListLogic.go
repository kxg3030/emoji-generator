package logic

import (
	"emoji/pkg/config"
	userEntity "emoji/pkg/model/entity"
	"emoji/pkg/unity"
	"github.com/gohouse/gorose"
)

type UserListLogic struct {
	orm *gorose.Session
}

var table string
func init()  {
	table = "user_list"
}

func NewUserListLogic(orm *gorose.Session)*UserListLogic  {
	return &UserListLogic{
		orm:orm,
	}
}

func (this *UserListLogic)FindUserRecord(openId string)map[string]interface{}  {
	result,err := this.orm.Table(table).Where("open_id",openId).First()
	unity.ErrorCheck(err)
	return result
}

func (this *UserListLogic)InsertUserRecord(user userEntity.UserList)bool  {
	id,err := this.orm.Table(table).Data(map[string]interface{}{
		"open_id"     : user.OpenId,
		"avatar"      : user.Avatar,
		"nick_name"   : user.NickName,
		"create_time" : unity.GetNowDateTime(config.SecondFormat),
	}).InsertGetId()
	if err != nil  {
		return false
	}
	return id >= 1
}

func (this *UserListLogic)UpdateUserColumn(openId string,user userEntity.UserList) bool {
	result,err := this.orm.Table(table).Where("open_id",openId).Data(map[string]interface{}{
		"avatar" : user.Avatar,
		"nick_name" : user.NickName,
	}).Update()
	if err != nil {
		return false
	}
	return result >= 1
}