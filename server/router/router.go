package router

import (
	"server/internal/user"
	"server/internal/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	api := r.Group("/api/v1/chat")
	{
		api.POST("/signup", userHandler.CreateUser)
		api.POST("/login", userHandler.Login)
		api.GET("/logout", userHandler.Logout)

		//websocket endpoints
		api.POST("/ws/createRoom", wsHandler.CreateRoom)
		api.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
		api.GET("/ws/getRooms", wsHandler.GetRooms)
		api.GET("/ws/getClients/:roomId", wsHandler.GetClients)
	}
}

func Start(addr string) error {
	return r.Run(addr)
}
