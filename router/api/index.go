package api

import (
	"myblog-server/router/api/userRouter"

	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {

	api := r.Group("/api")
	userRouter.UserInitRouter(api)
}
