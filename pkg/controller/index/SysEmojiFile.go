package index

import (
	"emoji/pkg/database"
	"emoji/pkg/model/logic"
	"emoji/pkg/system"
	"github.com/gin-gonic/gin"
	"html"
)

type SysEmojiFile struct {
	
}

func NewSysEmojiFile() *SysEmojiFile  {
	return &SysEmojiFile{
		
	}
}

func (this *SysEmojiFile)GetEmojiFileList(ctx *gin.Context)  {
	filed := "id,name,cover_url"
	page  := html.EscapeString(ctx.DefaultQuery("page","1"))
	size  := html.EscapeString(ctx.DefaultQuery("size","10"))
	result := logic.NewSysEmojiFileLogic(database.GetOrm()).SelectSysFileList(filed,page,size)
	if len(result) != 0 {
		system.PrintSuccess(ctx,203,"",result)
		return
	}
	system.PrintException(ctx,117,"",result)
}

func (this *SysEmojiFile)GetEmojiFileDetail(ctx *gin.Context)  {
	id := html.EscapeString(ctx.Query("id"))
	result := logic.NewSysEmojiFileLogic(database.GetOrm()).GetsSysFileFirstById(id)
	if len(result) >= 1 {
		system.PrintSuccess(ctx,205,"",result)
		return
	}
	system.PrintException(ctx,117,"",map[string]interface{}{})
}