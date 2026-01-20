package router

import (
	"os"
	"server/internal/user"
	"server/internal/ws"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	originString := os.Getenv("CLIENT_ORIGINS")
	var allowedOrigins []string

	if originString != "" {
		allowedOrigins = strings.Split(originString, ",")
	} else {
		allowedOrigins = []string{"http://localhost:5173"}
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, 
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api/v1/chat")
	{
		api.POST("/signup", userHandler.CreateUser)
		api.POST("/login", userHandler.Login)
		api.GET("/logout", userHandler.Logout)

		api.POST("/ws/createRoom", wsHandler.CreateRoom)
		api.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
		api.GET("/ws/getRooms", wsHandler.GetRooms)
		api.GET("/ws/getClients/:roomId", wsHandler.GetClients)
	}
}

func Start(addr string) error {
	return r.Run(addr)
}