package system

import "github.com/gin-gonic/gin"

var Exception = map[int]string{
	100 : "不支持的文件格式",
	101 : "上传文件失败",
	102 : "上传文件覆盖成功",
	103 : "媒体文件或模板文件不存在",
	104 : "更新文件url失败",
	105 : "更新用户信息失败",
	106 : "授权码过期",
	107 : "授权码非法",
	108 : "服务器内部错误",
	109 : "获取用户openId失败",
	110 : "参数不足",
	111 : "签名验证失败",
	112 : "文件不存在",
	113 : "请重新登陆",
	114 : "自定义语句数错误",
	115 : "模板文件不存在",
	116 : "媒体文件不存在",
	117 : "空数据",
	118 : "ssl验证失败",
	119 : "404",
}

var Success   = map[int]string{
	200 : "上传成功",
	201 : "创建成功",
	202 : "登陆成功",
	203 : "成功",
}


func  PrintSuccess(ctx *gin.Context,code int,msg string,data interface{}) {
	if val,ok := Success[code];ok && msg == ""{
		msg = val
	}

	result := map[string]interface{}{
		"status" : true,
		"msg"    : msg,
		"code"   : code,
		"data"   : data,
	}
	ctx.JSON(200,result)
}

func PrintException(ctx *gin.Context,code int,msg string,data interface{}) {
	if val,ok := Exception[code];ok && msg == ""{
		msg = val
	}

	result := map[string]interface{}{
		"status" : false,
		"msg"    : msg,
		"code"   : code,
		"data"   : data,
	}
	ctx.JSON(200,result)
}

func GetExceptionMessage(code int)map[string]interface{}  {
	var msg string
	if val,ok := Exception[code];ok && msg == ""{
		msg = val
	}
	result := map[string]interface{}{
		"status" : false,
		"msg"    : msg,
		"code"   : code,
		"data"   : map[string]interface{}{},
	}
	return result
}