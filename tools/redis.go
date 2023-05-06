package tools

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis/v8"
	"sxdbman/datam"
	// "time"
	"context"
)

func InitRedis() {
	//func InitRedis() (rds *redis.Client) {

	var RedisDb *redis.Client
	var conf datam.Redisconn
	if _, err := toml.DecodeFile("./conf/redis.toml", &conf); err != nil {
		// handle error
		return
	}

	fmt.Println(conf.Addr)
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     conf.Addr, // redis地址
		Password: "",        // redis密码，没有则留空
		DB:       0,         // 默认数据库，默认是0
	})

	ctx := context.Background()
	err := RedisDb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}
	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	//return RedisDb
}
