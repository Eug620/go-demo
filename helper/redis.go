package engine

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/joho/godotenv/autoload" // 自动加载env
	"gopkg.in/ini.v1"
)

type RedisHelper struct {
	*redis.Client
}

var redisHelper *RedisHelper

var redisOnce sync.Once

func GetRedisHelper() *RedisHelper {
	return redisHelper
}

func NewRedisHelper() *redis.Client {
	cfg, cfgErr := ini.Load("config/dev.ini")
	if cfgErr != nil {
		fmt.Printf("Fail to read file: %v", cfgErr)
		os.Exit(1)
	}

	fmt.Println("env: ")

	// 读取操作，默认分区可以使用空字符串表示
	// cfg.Section("").Key("name").String()
	// cfg.Section("user").Key("account").String()

	// 修改某个值然后进行保存
	// cfg.Section("").Key("name").SetValue("newName")
	// cfg.SaveTo("config/my.ini.local")
	// cfg.Section("redis").Key("account").SetValue("654321")

	Addr := cfg.Section("redis").Key("Addr").String()
	Password := os.Getenv("Password")

	rdb := redis.NewClient(&redis.Options{
		Addr:         Addr,
		Password:     Password,
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	redisOnce.Do(func() {
		rdh := new(RedisHelper)
		rdh.Client = rdb
		redisHelper = rdh
	})

	return rdb
}
