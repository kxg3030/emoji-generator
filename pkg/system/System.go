package system

import "github.com/gin-gonic/gin"

var Exception = map[int]map[string]string{
	100 : {
		"out" : "文件格式错误",
		"in"  : "不支持的文件格式",
	},
	101 : {
		"out" : "上传文件失败",
		"in"  : "创建文件模板错误",
	},
	103 : {
		"out" : "生成GIF失败",
		"in"  : "媒体文件或模板文件不存在",
	},
	104 : {
		"out" : "生成GIF失败",
		"in"  : "更新文件url失败",
	},
	105 : {
		"out" : "登录失败",
		"in"  : "添加用户信息失败",
	},
	106 : {
		"out" : "授权码过期",
		"in"  : "授权码过期",
	},
	107 : {
		"out" : "授权码非法",
		"in"  : "授权码非法",
	},
	108 : {
		"out" : "服务器内部错误",
		"in"  : "服务器内部错误",
	},
	109 : {
		"out" : "登录失败",
		"in"  : "获取用户openId失败",
	},
	110 : {
		"out" : "参数不足",
		"in"  : "参数code不存在",
	},
	111 : {
		"out" : "签名错误",
		"in"  : "签名验证失败",
	},
	112 : {
		"out" : "无法创建GIF",
		"in"  : "mp4或ass文件不存在",
	},
	401 : {
		"out" : "请重新登陆",
		"in"  : "用户openid丢失或jwt错误",
	},
	114 : {
		"out" : "上传失败",
		"in"  : "自定义语句数错误",
	},
	115 : {
		"out" : "创建GIF失败",
		"in"  : "ass模板文件不存在",
	},
	116 : {
		"out" : "创建GIF失败",
		"in"  : "mp4媒体文件不存在",
	},
	117 : {
		"out" : "空数据",
		"in"  : "数据为空",
	},
	118 : {
		"out" : "ssl验证失败",
		"in"  : "ssl验证失败",
	},
	119 : {
		"out" : "路由不存在",
		"in"  : "404 not found",
	},
	220 : {
		"out" : "创建失败",
		"in"  : "没有匹配任何字符串",
	},
	221 : {
		"out" : "登陆失败",
		"in"  : "更新用户信息失败",
	},
	222 : {
		"out" : "创建失败",
		"in"  : "执行ffmpeg命令失败",
	},
	223 : {
		"out" : "创建失败",
		"in"  : "插入用户创建记录失败",
	},
}

var Success   = map[int]map[string]string{
	200 : {
		"out" : "成功",
		"in"  : "添加文件信息成功",
	},
	201 : {
		"out" : "成功",
		"in"  : "管理员创建GIF成功",
	},
	202 : {
		"out" : "成功",
		"in"  : "更新用户信息成功",
	},
	203 : {
		"out" : "成功",
		"in"  : "用户获取GIF列表成功",
	},
	204 : {
		"out" : "成功",
		"in"  : "更新文件信息成功",
	},
	205 : {
		"out" : "成功",
		"in"  : "成功",
	},
}


func  PrintSuccess(ctx *gin.Context,code int,msg string,data interface{}) {
	if val,ok := Success[code];ok && msg == ""{
		msg = val["out"]
	}

	result := map[string]interface{}{
		"status" : true,
		"msg"    : msg,
		"code"   : 200,
		"data"   : data,
	}
	ctx.JSON(200,result)
}

func PrintException(ctx *gin.Context,code int,msg string,data interface{}) {
	if val,ok := Exception[code];ok && msg == ""{
		msg = val["out"]
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
		msg = val["out"]
	}
	result := map[string]interface{}{
		"status" : false,
		"msg"    : msg,
		"code"   : code,
		"data"   : map[string]interface{}{},
	}
	return result
}