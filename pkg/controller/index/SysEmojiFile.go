package index

import (
	"emoji/pkg/database"
	"emoji/pkg/model/logic"
	"emoji/pkg/system"
	"github.com/gin-gonic/gin"
)

type SysEmojiFile struct {
	
}

func NewSysEmojiFile() *SysEmojiFile  {
	return &SysEmojiFile{
		
	}
}

func (this *SysEmojiFile)GetEmojiFileList(ctx *gin.Context)  {
	filed := "id,name,cover_url,md5_encode"
	result := logic.NewSysEmojiFileLogic(database.GetOrm()).SelectSysFileList(filed)
	system.PrintSuccess(ctx,203,"",result)
}