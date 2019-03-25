package service

import (
	"emoji/pkg/model/service/entity"
	"emoji/pkg/unity"
	"encoding/json"
	"fmt"
)
const (
	MINI_APP_ID      = "wx5ea0c993a127d18f"
	MINI_SECRET      = "1ffddf30b8909a739413fa16b5c59f1e"
	MINI_PROGRAM_BASE= "https://api.weixin.qq.com"
	MINI_USER_OPEN_ID= "/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
)

type WeChat struct {

}

func NewWeChat()*WeChat  {
	return &WeChat{

	}
}

func (this *WeChat)GetUserOpenId(code string)entity.UserInfo  {
	var userInfo entity.UserInfo
	response := unity.HttpGet(fmt.Sprintf(MINI_PROGRAM_BASE + MINI_USER_OPEN_ID,MINI_APP_ID,MINI_SECRET,code))
	err := json.Unmarshal(response,&userInfo)
	unity.ErrorCheck(err)
	return userInfo
}