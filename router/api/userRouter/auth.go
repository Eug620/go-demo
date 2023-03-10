/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2023-01-31 15:57:46
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2023-02-01 11:07:55
 * @FilePath     : /go-demo/router/api/userRouter/auth.go
 * @Description  : filename
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package userRouter

import (
	"fmt"
	"log"
	engine "myblog-server/helper"
	UserMiddleWare "myblog-server/middleware/user"

	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	name string `json:"name"`

	account int64 `json:"account"`
}

var redisHelper *engine.RedisHelper

func UserAuthRouter(g *gin.RouterGroup) {
	redisHelper = engine.GetRedisHelper()

	g.GET("/login", func(ctx *gin.Context) {
		fmt.Println("env: ", os.Getenv("env"))

		// 获取全局中间件写入的值
		req, _ := ctx.Get("request")
		fmt.Println("request:", req)

		// 取值
		token := ctx.GetHeader("token") // 获取请求头参数

		// 测试redis
		data, _ := redisHelper.HGetAll("type_hash").Result()
		fmt.Println(data)

		ctx.JSON(200, gin.H{"msg": "login success!", "token": token, "data": data, "req": req})
	})

	g.GET("/redis/get", func(ctx *gin.Context) {
		data, _ := redisHelper.HGetAll("type_hash").Result()
		ctx.JSON(200, gin.H{"msg": "redis get success!", "data": data})

	})
	g.GET("/redis/set", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "张三")
		password := ctx.DefaultQuery("password", "22")
		_ = redisHelper.HSet("type_hash", name, password).Err()
		ctx.JSON(200, gin.H{"msg": "redis set success!"})
	})

	g.GET("/auth", UserMiddleWare.AuthMiddleWare(), func(ctx *gin.Context) {
		// 获取全局中间件写入的值
		req, _ := ctx.Get("request")
		fmt.Println("request:", req)
		// 获取局部中间件写入的值
		auth, _ := ctx.Get("auth")
		fmt.Println("request:", auth)

		ctx.JSON(200, gin.H{"msg": "auth!", "auth": auth, "req": req})
	})

	// name形式
	g.GET("/name/:account", func(ctx *gin.Context) {
		account := ctx.Param("account")
		ctx.JSON(200, gin.H{"name": account})
	})

	// action形式
	g.GET("/user/:account/*action", func(ctx *gin.Context) {
		account := ctx.Param("account")
		action := ctx.Param("action")

		ctx.JSON(200, gin.H{"name": account, "action": action})
	})

	// GET 参数
	g.GET("/query", func(ctx *gin.Context) {

		username := ctx.DefaultQuery("username", "Guy")
		account := ctx.Query("account") // 是 c.Request.URL.Query().Get("account") 的简写
		ctx.JSON(200, gin.H{"username": username, "account": account})
	})

	// POST 参数

	g.POST("/post/:types", func(ctx *gin.Context) {
		types := ctx.Param("types")

		if types == "1" {
			// 获取表单参数
			message := ctx.PostForm("userName")              // 表单参数
			nick := ctx.DefaultPostForm("account", "123456") // 此方法可以设置默认值，和上面的get一样
			ctx.JSON(200, gin.H{"message": message, "nick": nick, "types": "1"})

		} else if types == "2" {

			// 获取body中的参数方式一

			json := make(map[string]interface{}) //注意该结构接受的内容

			ctx.BindJSON(&json)

			log.Printf("%v", &json)
			ctx.JSON(200, gin.H{"json": &json, "types": "2"})

		} else if types == "3" {
			// TODO未走通
			// 获取body中的参数方式二
			json := User{}
			ctx.BindJSON(&json)
			log.Printf("%v", &json)
			ctx.JSON(200, gin.H{"json": &json, "types": "3"})
		} else {
			ctx.JSON(200, gin.H{"msg": "types in undefined"})
		}
	})

}
