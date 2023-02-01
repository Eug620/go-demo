package main

import (
	"fmt"
	"log"
	engine "myblog-server/helper"
	baseRouter "myblog-server/router"
)

func main() {

	// redis
	rdb := engine.NewRedisHelper()
	if _, err := rdb.Ping().Result(); err != nil {
		log.Fatal(err.Error())
		return
	}

	S := baseRouter.InitRouter()
	err := S.Run(":8080")
	if err != nil {
		fmt.Println("启动失败...")
	}

}
