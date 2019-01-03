package main

import (
	"github.com/astaxie/beego/cache"
	_ "myAdmin/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	"time"
)
func init(){
	cacheconn := beego.AppConfig.String("redis.conn")
	cachepwd := beego.AppConfig.String("redis.pwd")
	cachedb := beego.AppConfig.String("redis.cachedbname")
	cachestr := `{"key":"victor","conn":"`+cacheconn+`","dbNum":"`+cachedb+`","password":"`+ cachepwd+`"}`
	beego.Info(cachestr)
	bm,err :=cache.NewCache("redis",cachestr)
	if err != nil{
		beego.Info("redis作为缓存错误是:",err)
	}
	///beego.Info(bm)
	bm.Put("name",1,10*time.Second)
	bm.Put("age",1000,1000*time.Second)
}

func main() {
	beego.Run()
}

