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
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SysEmojiFile struct {

}

func NewSysEmojiFile()*SysEmojiFile  {
	return &SysEmojiFile{

	}
}

// upload file
func (this *SysEmojiFile)UploadFile(ctx *gin.Context)  {
	var emojiFile entity.EmojiFile
	filePtrUse,header,err := ctx.Request.FormFile("file")
	sentence := ctx.PostForm("sentence")
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
	templateFileOri   := unity.FileMakeGetPtr(templateFilePath)
	if templateFileOri != nil {
		var fileNewStr string
		regRule := "Dialogue: (\\d,\\d:\\d{0,2}:\\d{0,2}\\.\\d{0,2}){2},\\w+,(,\\d{0,2}){3}(,){2}<\\?loading-[%s]\\?>"
		emojiFile = entity.EmojiFile{
			Name      : fileNameOnly,
			Extension : fileExt,
			Path      : templateFilePath,
			Md5Encode : unity.Md5String(fileNameOnly),
			CreateTime: time.Now().Format(config.SecondFormat),
			BasePath  : templateFileDir,
		}
		_,err  = io.Copy(templateFileOri,filePtrUse)
		if(fileExt == ".ass"){
			templateFileUse   := templateFileDir + "_temp_" +fileName
			if unity.FileMake(templateFileDir + "_temp_" + fileName) == false{
				unity.ErrorCheck(errors.New("create user  temp .ass file failed"))
			}
			defer func() {
				unity.ErrorCheck(templateFileOri.Close())
			}()
			unity.ErrorCheck(err)
			fileStr,err := ioutil.ReadFile(templateFilePath)
			unity.ErrorCheck(err)
			fileNewStr     = string(fileStr)
			sentenceSlice := strings.Split(sentence,"|")
			matchLineCount,err := regexp.Compile(fmt.Sprintf(regRule,"\\d{0,3}"))
			matchLineStrCount := matchLineCount.FindAllString(string(fileStr),-1)
			if len(matchLineStrCount) != len(sentenceSlice){
				system.PrintException(ctx,114,"", map[string]interface{}{})
				return
			}
			for key,_ := range sentenceSlice{
				matchLineReg,err := regexp.Compile(fmt.Sprintf(regRule,strconv.Itoa(key)))
				unity.ErrorCheck(err)
				matchLineString := matchLineReg.FindString(string(fileNewStr))
				if len(matchLineString) == 0 {
					unity.ErrorCheck(errors.New("can not match anything at key " + strconv.Itoa(key)))
				}
				matchLinePartReg,err := regexp.Compile("Dialogue: (\\d,\\d:\\d{0,2}:\\d{0,2}\\.\\d{0,2}){2},\\w+,(,\\d{0,2}){3}(,){2}")
				unity.ErrorCheck(err)
				matchNewString := matchLinePartReg.FindString(matchLineString)
				matchNewString += sentenceSlice[key]
				fileNewStr = matchLineReg.ReplaceAllString(fileNewStr,matchNewString)
			}
			emojiFile.SentenceCount = len(matchLineStrCount)
			emojiFile.Sentence      = sentence
			err = ioutil.WriteFile(templateFileUse,[]byte(fileNewStr),os.ModePerm)
			unity.ErrorCheck(err)
		}
		logicInstance := logic.NewSysEmojiFileLogic(database.GetOrm())
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
func (this *SysEmojiFile)GeneratorGifFromVideo(ctx *gin.Context)  {
	var emoji entity.EmojiFile
	var sentenceCount int64
	code := ctx.Param("encode")
	emoji = entity.EmojiFile{
		Md5Encode : code,
	}
	result := logic.NewSysEmojiFileLogic(database.GetOrm()).GetSysFileList(emoji)
	var sysFileName string
	gifExt := ".gif"
	pngExt := ".png"
	if len(result) == 2 {
		var sysFilePath,sysAssFile,sysSaveName string
		for _,val := range result{
			if val["extension"].(string) == ".mp4"{
				sysFilePath = val["path"].(string)
			}
			if val["extension"].(string) == ".ass"{
				sysAssFile  = val["base_path"].(string)
				if val,ok := val["sentence_count"].(int64);ok{
					sentenceCount = val
				}
			}

			sysSaveName = val["name"].(string)
			sysFileName = val["name"].(string)
		}

		sysAssFile += "_temp_" + sysSaveName + ".ass"
		sysSaveName = config.ASSETS_PATH + "system/" + sysSaveName + gifExt
		protocol,_ := ctx.Get("protocol")
		imageUrl   := protocol.(string)  + ctx.Request.Host   + "/assets/system/" + sysFileName + gifExt
		coverUrl   := protocol.(string)  + ctx.Request.Host   + "/assets/system/" + sysFileName + pngExt
		coverSave  := config.ASSETS_PATH + "system/" + sysFileName + pngExt
		if unity.DirExistValidate(config.ASSETS_PATH + "system/") == false{
			unity.DirMakeAll(config.ASSETS_PATH + "system/")
		}
		if unity.DirExistValidate(sysAssFile) == false{
			system.PrintException(ctx,115,"",map[string]interface{}{})
			return
		}
		if unity.DirExistValidate(sysFilePath) == false{
			system.PrintException(ctx,116,"",map[string]interface{}{})
			return
		}
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
		if logic.NewSysEmojiFileLogic(database.GetOrm()).UpdateSysFileImageUrl(imageUrl,coverUrl,code,sentenceCount){
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
func (this SysEmojiFile)GeneratorCoverFromVideo(sysFilePath string,sysFileName string,saveFilePath string,code string)bool  {
	var command =  &exec.Cmd{}
	command = exec.Command("ffmpeg","-y",
		"-i", sysFilePath, "-vframes", "1", "-ss", "0:0:0", "-an",
		"-vcodec", "png", "-f", "rawvideo", "-s", "100*100", saveFilePath)
	if _,err := command.CombinedOutput();err != nil {
		return false
	}
	return true
}





