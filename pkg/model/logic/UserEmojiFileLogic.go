package logic

import (
	"emoji/pkg/model/entity"
	"github.com/gohouse/gorose"
)

type UserEmojiFileLogic struct {
	orm *gorose.Session
}

func NewUserEmojiFileLogic(orm *gorose.Session)*UserEmojiFileLogic  {
	return &UserEmojiFileLogic{
		orm:orm,
	}
}

func (this *UserEmojiFileLogic)InsertNewRecord(emoji entity.UserEmojiFile)bool  {
	result,err := this.orm.Table("user_emoji_file").Data(map[string]interface{}{
		"open_id" : emoji.OpenId,
		"image_url" : emoji.ImageUrl,
		"create_time" : emoji.CreateTime,
	}).InsertGetId()
	if err != nil {
		return false
	}
	return result >= 1
}
