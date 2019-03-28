package bootstrap

import (
	"emoji/pkg/config"
	"emoji/pkg/database"
	"emoji/pkg/logger"
	"emoji/pkg/middleware"
	"emoji/pkg/router"
	"emoji/pkg/task"
	"emoji/pkg/unity"
	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/http"
)


type Bootstrap struct {
	Framework  *gin.Engine
}


func NewBootstrap(framework *gin.Engine)*Bootstrap  {
	return &Bootstrap{
		Framework:framework,
	}
}

func (this *Bootstrap)Init()*Bootstrap  {
	this.initLoggerFramework()
	this.setGlobalMiddleware()
	this.setDebugMode()
	this.setAssetsPath()
	this.initFrameworkRouter()
	this.initEnv()
	this.initTask()
	this.setOrm()
	return this
}

func (this *Bootstrap)Run(port string)  {
	unity.ErrorCheck(this.Framework.Run(port))
}

func (this *Bootstrap)initFrameworkRouter()*Bootstrap  {
	this.Framework = router.NewRouter(this.Framework).RegisterRouter()
	return this
}

func (this *Bootstrap)setDebugMode()  {
	debugMode := config.Config["DebugMode"].(bool)
	if debugMode{
		this.Framework.Use(gin.Logger())
	}
}

func (this *Bootstrap)setGlobalMiddleware()  {
	middlewareHandle,ok := config.Config["GlobalMiddleWare"].([]middleware.MiddlewareInterface)
	if ok{
		for _,val := range middlewareHandle{
			this.Framework.Use(val.Render())
		}
	}
}

func (this *Bootstrap)setAssetsPath()  {
	this.Framework.StaticFS("/assets",http.Dir(config.ASSETS_PATH))
}

func (this *Bootstrap)setOrm()  {
	var err error
	config.Database.Dsn   = unity.GetEnvVal("mysql")
	database.Database,err = gorose.Open(config.Database)
	unity.ErrorCheck(err)
}

func (this *Bootstrap)initEnv()  {
	unity.ErrorCheck(godotenv.Load())
}

func (this *Bootstrap)initLoggerFramework()  {
	hook := lumberjack.Logger{
		Filename  : config.LOG_PATH + unity.GetNowDateTime(config.DateFormat), // 日志文件路径
		MaxSize   : 128,                      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                       // 日志文件最多保存多少个备份
		MaxAge    :  7,                       // 文件最多保存多少天
		Compress  : true,                     // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey      :  "time",
		LevelKey     :  "level",
		NameKey      :  "logger",
		CallerKey    :  "line",
		MessageKey   :  "msg",
		StacktraceKey:  "stacktrace",
		LineEnding   :  zapcore.DefaultLineEnding,
		EncodeLevel  :  zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime   :  logger.EncodeTime,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller  : zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName   :  zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                    // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),      // 打印到控制台和文件
		atomicLevel,                                              // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", "emoji"))
	// 构造日志
	logger.Logger = zap.New(core, caller, development, filed)
}

func (this Bootstrap)initTask()  {
	task.NewTask(config.RUNTIME_PATH).DeleteExpireAssFile()
}




