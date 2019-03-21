package unity

import (
	"crypto/md5"
	"emoji/pkg/logger"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"
)

func ErrorCheck(err error)  {
	if err != nil {
		logger.Logger.Error(err.Error())
		panic(err.Error())
	}
}

func Md5String(str string) string  {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return fmt.Sprintf("%x",ctx.Sum(nil))
}

func GetNowDateTime(format string) string  {
	datetime := time.Now().Format(format)
	return datetime
}

func DirExistValidate(dir string)bool  {
	_,err := os.Stat(dir)
	if os.IsNotExist(err){
		return false
	}
	if err != nil {
		return false
	}
	return true
}

func DirMakeAll(dir string) bool  {
	err := os.MkdirAll(dir,os.ModePerm)
	ErrorCheck(err)
	return true
}

func FileMake(fileName string)  {
	_,err := os.Open(fileName)
	if err != nil && os.IsNotExist(err){
		fi,err := os.Create(fileName)
		ErrorCheck(err)
		defer func() {
			ErrorCheck(fi.Close())
		}()
	}
}

func FileMakeGetPtr(fileName string)*os.File  {
	var fi *os.File
	var err error
	fi,err = os.OpenFile(fileName,os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil && os.IsNotExist(err){
		fi,err = os.Create(fileName)
		ErrorCheck(err)
		return fi
	}
	return fi
}


func DynamicProxyCall(object interface{},method string,args ...interface{})[]reflect.Value  {
	relObj := reflect.ValueOf(object)
	if relObj.Type().NumIn() != len(args){
		ErrorCheck(errors.New("param count is not match"))
	}
	params := make([]reflect.Value,len(args))
	for key,val := range args{
		params[key] = reflect.ValueOf(val)
	}
	return reflect.ValueOf(object).MethodByName(method).Call(params)
}

func GetEnvVal(key string)string  {
	return os.Getenv(key)
}
