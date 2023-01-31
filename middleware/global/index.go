/*
 * @Author       : eug yyh3531@163.com
 * @Date         : 2023-01-31 17:33:21
 * @LastEditors  : eug yyh3531@163.com
 * @LastEditTime : 2023-01-31 17:47:09
 * @FilePath     : /go-server/middleware/global/auth.go
 * @Description  : filename
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package AuthMiddleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 全局中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()

		fmt.Println("全局中间件开始执行了")

		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "全局中间件")

		status := c.Writer.Status()

		fmt.Println("全局中间件执行完毕", status)

		t2 := time.Since(t)

		fmt.Println("time:", t2)

	}
}
