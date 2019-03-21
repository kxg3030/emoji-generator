package logic

import (
	"emoji/pkg/model/entity"
	"emoji/pkg/unity"
	"github.com/gohouse/gorose"
)

type EmojiFileLogic struct {
	Orm  *gorose.Session
}

func NewEmojiFileLogic(orm *gorose.Session)*EmojiFileLogic  {
	return &EmojiFileLogic{
		Orm : orm,
	}
}

func (this *EmojiFileLogic) InsertNewFileRecord(emoji entity.EmojiFile) bool  {
	result ,err  := this.Orm.Table("sys_emoji_file").Where(map[string]interface{}{
		"md5_encode"  : emoji.Md5Encode,
		"extension"   : emoji.Extension,
	}).First()
	if len(result) >= 1{
		return false
	}
	insertId,err := this.Orm.Table("sys_emoji_file").Data(map[string]interface{}{
		"name"        : emoji.Name,
		"path"        : emoji.Path,
		"extension"   : emoji.Extension,
		"md5_encode"  : emoji.Md5Encode,
		"create_time" : emoji.CreateTime,
		"base_path"   : emoji.BasePath,
		"sentence"    : emoji.Sentence,
		"sentence_count"    : emoji.SentenceCount,
	}).InsertGetId()
	unity.ErrorCheck(err)
	return insertId >= 1
}

func (this *EmojiFileLogic)GetSysFileList(emoji entity.EmojiFile)[]map[string]interface{}  {
	result,err := this.Orm.Table("sys_emoji_file").Where(map[string]interface{}{
		"md5_encode" : emoji.Md5Encode,
	}).Fields("path,base_path,extension,name").Get()
	unity.ErrorCheck(err)
	return result
}

func (this *EmojiFileLogic)UpdateSysFileImageUrl(url string,cover string,md5 string)bool  {
	_,err := this.Orm.Table("sys_emoji_file").Where("md5_encode",md5).Data(map[string]interface{}{
		"image_url" : url,
		"cover_url" : cover,
	}).Update()
	unity.ErrorCheck(err)
	return true
}