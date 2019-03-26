package system

import "github.com/gin-gonic/gin"

var Exception = map[int]string{
	100 : "upload file ext is not support",
	101 : "upload file failed",
	102 : "insert record to database failed",
	103 : "audio  file count not equal two,can not convert to gif",
	104 : "update url failed",
	105 : "update user failed",
	106 : "token expire",
	107 : "token illegal",
	108 : "inner error",
	109 : "get user open id failed",
	110 : "param required",
	111 : "signature illegal",
}

var Success   = map[int]string{
	200 : "upload  success",
	201 : "convert success",
	202 : "login success",
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
	if val,ok := Exception[code];ok{
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