package admin

import (
	"emoji/pkg/config"
	"emoji/pkg/database"
	"emoji/pkg/model/entity"
	"emoji/pkg/model/logic"
	"emoji/pkg/system"
	"emoji/pkg/unity"
	"emoji/pkg/validate"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os/exec"
	"path"
	"strings"
	"time"
)

type Emoji struct {

}

func NewEmoji()*Emoji  {
	return &Emoji{

	}
}

func (this *Emoji)UploadFile(ctx *gin.Context)  {
	var emojiFile entity.EmojiFile
	filePtr,header,err := ctx.Request.FormFile("file")
	unity.ErrorCheck(err)
	fileName := header.Filename
	fileExt  := path.Ext(fileName)
	if validate.ExtIsIllegal(fileExt) == false{
		system.PrintException(ctx,100,"", map[string]interface{}{})
		return
	}
	fileNameOnly := strings.TrimSuffix(fileName,fileExt)
	templateFileDir  := config.TEMPLATE_PATH + unity.Md5String(fileNameOnly) + config.OS_SEPREATOR
	if unity.DirExistValidate(templateFileDir) == false{
		unity.DirMakeAll(templateFileDir)
	}
	templateFilePath  := templateFileDir + fileName
	templateFilePtr   := unity.FileMakeGetPtr(templateFilePath)
	defer func() {
		unity.ErrorCheck(templateFilePtr.Close())
	}()
	if templateFilePtr != nil {
		emojiFile = entity.EmojiFile{
			Name      : fileNameOnly,
			Extension : fileExt,
			Path      : templateFilePath,
			Md5Encode : unity.Md5String(fileNameOnly),
			CreateTime: time.Now().Format(config.SecondFormat),
			BasePath  : templateFileDir,
		}
		_,err := io.Copy(templateFilePtr,filePtr)
		unity.ErrorCheck(err)
		logicInstance := logic.NewEmojiFileLogic(database.GetOrm())
		if ok := logicInstance.InsertNewFileRecord(emojiFile);ok{
			system.PrintSuccess(ctx,200,"", map[string]interface{}{})
			return
		}
		system.PrintException(ctx,102,"", map[string]interface{}{})
		return
	}
	system.PrintException(ctx,101,"", map[string]interface{}{})
}

func (this *Emoji)GeneratorGifFromVideo(ctx *gin.Context)  {
	var emoji entity.EmojiFile
	code := ctx.Param("code")
	emoji = entity.EmojiFile{
		Md5Encode : code,
	}
	result := logic.NewEmojiFileLogic(database.GetOrm()).GetSysFileList(emoji)
	var sysFileName string
	if len(result) == 2 {
		var sysFilePath,sysAssFile,sysSaveName string
		for _,val := range result{
			if ok := val["extension"].(string) == ".mp4";ok{
				sysFilePath = val["path"].(string)
			}
			if ok := val["extension"].(string) == ".ass";ok{
				sysAssFile  = val["path"].(string)
			}
			sysSaveName = val["name"].(string)
			sysFileName = val["name"].(string)
		}
		sysSaveName = config.ASSETS_PATH + sysSaveName + ".gif"
		var command =  &exec.Cmd{}
		command = exec.Command("ffmpeg","-y","-i",sysFilePath,"-vf",fmt.Sprintf("ass=%s",sysAssFile),sysSaveName)
		if _,err := command.CombinedOutput();err != nil{
			unity.ErrorCheck(err)
		}
		imageUrl := ctx.Request.Host + "/assets/" + sysFileName + ".gif"
		if logic.NewEmojiFileLogic(database.GetOrm()).UpdateSysFileImageUrl(imageUrl,code){
			system.PrintSuccess(ctx,201,"",map[string]interface{}{
				"url" : imageUrl,
			})
			return
		}
		system.PrintException(ctx,104,"",map[string]interface{}{})
		return
	}
	system.PrintException(ctx,103,"",map[string]interface{}{})
}





