package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/beego/cache"
	"time"
)

var redisClient cache.Cache
var redisTime, _ = beego.AppConfig.Int("redisTime")

// 初始化redis
func init() {
	config := map[string]string{
		"key":      beego.AppConfig.String("redisKey"),
		"conn":     beego.AppConfig.String("redisConn"),
		"dbNum":    beego.AppConfig.String("redisDbNum"),
		"password": beego.AppConfig.String("redisPwd"),
	}
	bytes, _ := json.Marshal(config)
	redisClient, err = cache.NewCache("redis", string(bytes))
	if err != nil {
		/*controllers.ExecError("init redis db failed", err, controllers.DBError, controllers.InitError)*/
	} else {
		fmt.Println("redis connect successful")
	}
}

type cacheDb struct{}

var CacheDb = &cacheDb{}

func (c cacheDb) Set(key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	redisClient.Put(key, string(bytes), time.Second*time.Duration(redisTime))
}
func (c cacheDb) Get(key string, obj interface{}) bool {
	if redisStr := redisClient.Get(key); redisStr != nil {
		fmt.Println("Redis is reading..." + key + ":")
		redisValue, ok := redisStr.([]uint8)
		if !ok {
			/*controllers.ExecError("getRedisFailed", nil, controllers.DBError)*/
			return false
		}
		json.Unmarshal([]byte(redisValue), obj)
		return true
	}
	return false
}
