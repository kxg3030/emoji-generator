package index

import (
	"emoji/pkg/config"
	"emoji/pkg/unity"
	"emoji/pkg/validate"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"html"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strings"
)

type Emoji struct {
	// 用户IP
	userHost        string
	// 用户Ass文件名
	userAssFileName string
	// 用户Ass存储路径
	userAssFilePath string
	// 用户生成文件存储路径
	userFileSave    string
	// 用户自定义内容
	userAssSentence []string
	// 用户选期望转换的扩展名
	userFileExtension string
	// 系统多媒体文件路径
	sysFilePath     string
	// 系统Ass文件路径
	sysAssPath      string
}

func NewEmoji()*Emoji  {
	return &Emoji{}
}

func (this *Emoji) EmojiGenerator(ctx *gin.Context){
	userFileName:= html.EscapeString(ctx.Query("fileName"))
	assSentence := html.EscapeString(ctx.Query("sentence"))
	userUniqueId:= ctx.Request.Host
	this.userHost          = userUniqueId
	this.userAssSentence   = strings.Split(assSentence,",")
	this.userFileExtension = html.EscapeString(ctx.Query("extension"))
	this.userAssFileName   = userFileName
	this.userAssFilePath   = this.AnalysisAss(userFileName,assSentence,userUniqueId)
	this.ExecuteCommand()
}

func (this *Emoji)AnalysisAss(fileName string,sentence string,userId string)string  {
	var sentenceNew string
	if validate.ExtIsIllegal(this.userFileExtension) == false{
		unity.ErrorCheck(errors.New("ext not support"))
	}
	sysFileBase := unity.Md5String(fileName) + config.OS_SEPREATOR
	sysAssFile  := config.TEMPLATE_PATH + sysFileBase +fileName + config.ASS_FILE_EXT
	sysAudio    := config.TEMPLATE_PATH + sysFileBase +fileName
	this.sysAssPath = sysAssFile
	this.sysFilePath= sysAudio
	fileStr,err := ioutil.ReadFile(sysAssFile)
	unity.ErrorCheck(err)
	match,err := regexp.Compile("<\\?loading-[0-9]\\?>")
	unity.ErrorCheck(err)
	matchString := match.FindAllString(string(fileStr),-1)
	if len(matchString) == 0 {
		unity.ErrorCheck(errors.New("can not match anything"))
	}
	matchCount := len(matchString)
	sentenceStr:= strings.Split(sentence,"|")
	if matchCount != len(sentenceStr){
		unity.ErrorCheck(errors.New("params length error"))
	}
	userNewFile := config.RUNTIME_PATH + unity.GetNowDateTime(config.HourFormat) + config.OS_SEPREATOR
	this.userFileSave = userNewFile
	if unity.DirExistValidate(userNewFile) == false{
		unity.DirMakeAll(userNewFile)
	}
	userNewFile += "xx"
	userNewFile += config.ASS_FILE_EXT
	unity.FileMake(userNewFile)
	sentenceNew = string(fileStr)
	for key,val := range matchString{
		sentenceNew = strings.Replace(sentenceNew,val,sentenceStr[key],-1)
		err := ioutil.WriteFile(userNewFile,[]byte(sentenceNew),0)
		unity.ErrorCheck(err)
	}
	return userNewFile
}

func (this *Emoji)ExecuteCommand()  {
	var command =  &exec.Cmd{}
	sysFilePath := this.sysFilePath + ".mp4"
	usrAssFile  := this.userAssFilePath
	usrSavePath := this.userFileSave
	usrSavePath += "xx." +this.userFileExtension
	switch this.userFileExtension {
	case "mp4":

		break
	case "gif":
		command = exec.Command("ffmpeg","-y","-i",sysFilePath,"-vf",fmt.Sprintf("ass=%s",usrAssFile),usrSavePath)
		break
	case "png":
		break
	case "jpg":
		break
	case "jpeg":
		break
	default:
		unity.ErrorCheck(errors.New("file not support"))
	}
	if _,err := command.CombinedOutput();err != nil{
		unity.ErrorCheck(err)
	}
}
