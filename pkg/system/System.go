package system

import "github.com/gin-gonic/gin"

var Exception = map[int]string{
	100 : "upload file ext is not support",
	101 : "upload file failed",
	102 : "upload and override success",
	103 : "audio  file count not equal two,can not convert to gif",
	104 : "update url failed",
	105 : "update user failed",
	106 : "authorization expire",
	107 : "authorization illegal",
	108 : "inner error",
	109 : "get user open id failed",
	110 : "param required",
	111 : "signature illegal",
	112 : "file not found",
	113 : "please login ",
	114 : "sentence count cant not match ass",
	115 : ".ass file not exist",
	116 : ".mp4 file not exist",
	117 : "data is empty",
	118 : "ssl error",
	119 : "404 not found",
}

var Success   = map[int]string{
	200 : "upload  success",
	201 : "convert success",
	202 : "login success",
	203 : "success",
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