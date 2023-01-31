package baseRouter

import (
	AuthMiddleware "myblog-server/middleware/global"
	"myblog-server/router/api"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() *gin.Engine {

	r := gin.New()

	// 全局中间件
	r.Use(AuthMiddleware.MiddleWare())

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	api.InitApi(r)

	return r

}
