package main

import (
	"fmt"
	baseRouter "myblog-server/router"
	"os"

	"gopkg.in/ini.v1"
)

func main() {
	// env读取
	fmt.Println("env: ", os.Getenv("env"))

	// ini文件操作
	// 读取
	cfg, cfgErr := ini.Load("my.ini")
	if cfgErr != nil {
		fmt.Printf("Fail to read file: %v", cfgErr)
		os.Exit(1)
	}

	// 读取操作，默认分区可以使用空字符串表示
	cfg.Section("").Key("name").String()
	cfg.Section("user").Key("account").String()

	// 修改某个值然后进行保存
	cfg.Section("").Key("name").SetValue("newName")
	cfg.Section("user").Key("account").SetValue("654321")
	cfg.SaveTo("config/my.ini.local")

	// S := gin.Default()
	// S.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{"msg": "hello world!"})
	// })
	// S.GET("/user/:name/*action", func(c *gin.Context) {

	// 	name := c.Param("name") //通过Context的Param方法来获取API参数

	// 	action := c.Param("action")

	// 	fmt.Println(name, action)

	// 	//截取

	// 	action = strings.Trim(action, "/")

	// 	c.String(http.StatusOK, name+" is "+action)

	// })
	S := baseRouter.InitRouter()
	err := S.Run(":8080")
	if err != nil {
		fmt.Println("启动失败...")
	}

}
