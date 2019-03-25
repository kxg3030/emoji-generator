package admin

import (
	"emoji/pkg/config"
	"emoji/pkg/database"
	"emoji/pkg/model/entity"
	"emoji/pkg/model/logic"
	"emoji/pkg/system"
	"emoji/pkg/unity"
	"emoji/pkg/validate"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"time"
)

type EmojiFile struct {

}

func NewEmoji()*EmojiFile  {
	return &EmojiFile{

	}
}

// upload file
func (this *EmojiFile)UploadFile(ctx *gin.Context)  {
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
		if(fileExt != ".mp4"){
			fileStr,err := ioutil.ReadFile(templateFilePath)
			unity.ErrorCheck(err)
			match,err := regexp.Compile("Dialogue: (\\d,\\d:\\d{0,2}:\\d{0,2}\\.\\d{0,2}){2},\\w+,(,\\d{0,2}){3}(,){2}[\u4e00-\u9fa5a-zA-Z0-9]{0,}")
			unity.ErrorCheck(err)
			matchString := match.FindAllString(string(fileStr),-1)
			if len(matchString) == 0 {
				unity.ErrorCheck(errors.New("can not match anything"))
			}
			matchCount := len(matchString)
			for key,val := range matchString{
				match,err := regexp.Compile("Dialogue: (\\d,\\d:\\d{0,2}:\\d{0,2}\\.\\d{0,2}){2},\\w+,(,\\d{0,2}){3}(,){2}")
				matchStr  := match.FindString(val)
				unity.ErrorCheck(err)
				matchString[key] = strings.TrimPrefix(val,matchStr)
			}
			emojiFile.SentenceCount = matchCount
			emojiFile.Sentence      = strings.Join(matchString,"|")
		}
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

// async convert mp4 to gif
func (this *EmojiFile)GeneratorGifFromVideo(ctx *gin.Context)  {
	var emoji entity.EmojiFile
	code := ctx.Param("code")
	emoji = entity.EmojiFile{
		Md5Encode : code,
	}
	result := logic.NewEmojiFileLogic(database.GetOrm()).GetSysFileList(emoji)
	var sysFileName string
	gifExt := ".gif"
	pngExt := ".png"
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
		sysSaveName = config.ASSETS_PATH + sysSaveName + gifExt
		imageUrl := ctx.Request.Host + "/assets/" + sysFileName + gifExt
		coverUrl := ctx.Request.Host + "/assets/" + sysFileName + pngExt
		coverSave:= config.ASSETS_PATH + sysFileName + pngExt
		go func() {
			var command =  &exec.Cmd{}
			command = exec.Command("ffmpeg","-y","-i",sysFilePath,"-vf",fmt.Sprintf("ass=%s",sysAssFile),sysSaveName)
			if _,err := command.CombinedOutput();err != nil{
				unity.ErrorCheck(err)
			}
		}()
		go func() {
			this.GeneratorCoverFromVideo(sysFilePath,sysFileName,coverSave,code)
		}()
		if logic.NewEmojiFileLogic(database.GetOrm()).UpdateSysFileImageUrl(imageUrl,coverUrl,code){
			system.PrintSuccess(ctx,201,"",map[string]interface{}{
				"url" : imageUrl,
				"cov" : coverUrl,
			})
			return
		}
		system.PrintException(ctx,104,"",map[string]interface{}{})
		return
	}
	system.PrintException(ctx,103,"",map[string]interface{}{})
}

// async convert mp4 to png
func (this EmojiFile)GeneratorCoverFromVideo(sysFilePath string,sysFileName string,saveFilePath string,code string)bool  {
	var command =  &exec.Cmd{}
	command = exec.Command("ffmpeg","-y",
		"-i", sysFilePath, "-vframes", "1", "-ss", "0:0:0", "-an",
		"-vcodec", "png", "-f", "rawvideo", "-s", "100*100", saveFilePath)
	if _,err := command.CombinedOutput();err != nil {
		return true
	}
	return false
}





