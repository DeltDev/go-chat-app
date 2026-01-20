package main

import (
	"log"
	"os"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	userRep := user.NewRepository(dbConn.GetDB())

	userSvc := user.NewService(userRep, jwtSecret) 
	userHandler := user.NewHandler(userSvc, jwtSecret)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)

	addr := ":8080"
	if port := os.Getenv("SERVER_PORT"); port != "" {
		addr = port
	}

	log.Printf("Server starting on %s", addr)
	router.Start(addr)
}