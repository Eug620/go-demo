package userRouter

import (
	"github.com/gin-gonic/gin"
)

// 初始化模块路由

func UserInitRouter(r *gin.RouterGroup) {

	userAuth := r.Group("/user-auth")

	UserAuthRouter(userAuth)

}
