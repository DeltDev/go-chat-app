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
	// 1. Load .env file (optional, mainly for local dev)
    // In Docker, these will come from the -e flags
	_ = godotenv.Load()

	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	// 2. Get JWT Secret from Env
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	// 3. Initialize Repos and Services
	userRep := user.NewRepository(dbConn.GetDB())
    
    // Pass the jwtSecret to the service (Ensure your user_service.go accepts this!)
	userSvc := user.NewService(userRep, jwtSecret) 
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)

	// 4. Set Port
    // We default to :8080 inside the container. 
    // Docker will map external 15015 -> internal 8080.
	addr := ":8080"
	if port := os.Getenv("SERVER_PORT"); port != "" {
		addr = port
	}

	log.Printf("Server starting on %s", addr)
	router.Start(addr)
}