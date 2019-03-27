package unity

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
)

func ErrorCheck(err error)  {
	if err != nil {
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

func FileMake(fileName string)bool  {
	_,err := os.Open(fileName)
	if os.IsNotExist(err){
		fi,err := os.Create(fileName)
		ErrorCheck(err)
		defer func() {
			ErrorCheck(fi.Close())
		}()
		return  true
	}
	if err != nil {
		return false
	}
	return true
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

func HttpPost(url string,data string)string  {
	client   := &http.Client{}
	postData := bytes.NewBuffer([]byte(data))
	request,_:= http.NewRequest("POST",url,postData)
	request.Header.Set("Content-type", "application/json")
	response,_ := client.Do(request)
	if response.StatusCode == 200 {
		body,_ := ioutil.ReadAll(response.Body)
		return string(body)
	}
	return ""
}

func HttpGet(url string) []byte  {
	client     := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Connection", "keep-alive")
	response, _:= client.Do(request)
	if response.StatusCode == 200 {
		body, _:= ioutil.ReadAll(response.Body)
		return body
	}
	return []byte("")
}

func GetToken(key string,val string) string  {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(240)).Unix()
	claims["iat"] = time.Now().Unix()
	claims[key]   = val
	token.Claims  = claims
	tokenString, err := token.SignedString([]byte("emoji"))
	if err != nil {
		return ""
	}
	return tokenString
}
