package config

import (
	"emoji/pkg/middleware"
	"github.com/gohouse/gorose"
	_ "github.com/gohouse/gorose/driver/mysql"
)
const OS_SEPREATOR  = "/"
const ROOT_PATH     = "." + OS_SEPREATOR
const PAKAGE_PATH   = ROOT_PATH   + "pkg"      + OS_SEPREATOR
const ASSETS_PATH   = PAKAGE_PATH + "assets"   + OS_SEPREATOR
const TEMPLATE_PATH = PAKAGE_PATH + "template" + OS_SEPREATOR
const RUNTIME_PATH  = PAKAGE_PATH + "runtime"  + OS_SEPREATOR
const LOG_PATH      = PAKAGE_PATH + "logger"  + OS_SEPREATOR
const ASS_FILE_EXT  = ".ass"
const YearFormat    = "2006"
const MonthFormat   = "2006" + OS_SEPREATOR + "01"
const DateFormat    = "2006" + OS_SEPREATOR + "01" + OS_SEPREATOR + "02"
const HourFormat    = "2006" + OS_SEPREATOR + "01" + OS_SEPREATOR + "02" + OS_SEPREATOR + "15"
const MiniFormat    = "2006" + OS_SEPREATOR + "01" + OS_SEPREATOR + "02" + OS_SEPREATOR + "15" + OS_SEPREATOR + "04"
const SecondFormat  = "2006" + OS_SEPREATOR + "01" + OS_SEPREATOR + "02" + OS_SEPREATOR + "15" + OS_SEPREATOR + "04" + OS_SEPREATOR + "05"


var Extension = []string{
	".mp4",".gif",".png",".jpeg",".jpg",".ass",".srt",
}

var Config = map[string]interface{}{
	"ListenPort" : ":8080",
	"DebugMode"  : true,
	"MiddleWare" : []middleware.MiddlewareInterface{
		middleware.NewRouterMiddleware(),
		middleware.NewCrossSiteMiddleware(),
		middleware.NewRecoverMiddleware(),
	},
}

var Database = &gorose.DbConfigSingle{
	Driver:          "mysql",
	EnableQueryLog:  true,
	SetMaxOpenConns: 0,
	SetMaxIdleConns: 0,
	Prefix:          "xm_",
	Dsn:             "root:root@tcp(127.0.0.1:3306)/xm_emoji?charset=utf8",
}