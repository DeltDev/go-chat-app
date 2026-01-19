package router

import (
	"server/internal/user"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()

	api := r.Group("/api/v1/chat")
	{
		api.POST("/signup", userHandler.CreateUser)
		api.POST("/login", userHandler.Login)
		api.POST("/logout", userHandler.Logout)
	}
}

func Start(addr string) error {
	return r.Run(addr)
}
