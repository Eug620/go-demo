package UserMiddleWare

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义局部中间件
func AuthMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {

		t := time.Now()

		fmt.Println("局部中间件开始执行了")

		// 设置变量到Context的key中，可以通过Get()取
		c.Set("auth", "局部中间件")
		// 执行函数
		c.Next()

		// 中间件执行完后续的一些事情
		status := c.Writer.Status()

		fmt.Println("局部中间件执行完毕", status)

		t2 := time.Since(t)

		fmt.Println("time:", t2)

	}

}
