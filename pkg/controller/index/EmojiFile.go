package index

import (
	"emoji/pkg/database"
	"emoji/pkg/model/logic"
	"emoji/pkg/system"
	"github.com/gin-gonic/gin"
)

type EmojiFile struct {
	
}

func NewEmojiFile() *EmojiFile  {
	return &EmojiFile{
		
	}
}

func (this *EmojiFile)GetEmojiFileList(ctx *gin.Context)  {
	filed := "id,name,cover_url,md5_encode"
	result := logic.NewEmojiFileLogic(database.GetOrm()).SelectSysFileList(filed)
	system.PrintSuccess(ctx,200,"",result)
}