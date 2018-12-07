package utils
import (
	"bytes"
	"encoding/gob"
	"error"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	)

var cc cache.Cache

func InitCache() {
	host := beego.AppConfig.String("cache::redis_host")
	var err error
	defer func(){
		if r := recover(); r != nil {
			cc = nil
		}
	}()

	cc , err = cache.NewCache("redis",`{"conn":"`+host+`"}`)
	if err != nil {
		LogError("Connect tothe redis host" + host + "failed")
		LogError(err)
	}
}

func SetCache(key string,value interface{},timeout int) error {
	data, err := Encode(value)
	// Code ....
}	