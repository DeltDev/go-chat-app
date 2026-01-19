package main

import (
	"log"
	"os"
	"server/db"
	"server/internal/user"
	"server/router"
	"server/internal/ws"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)
	
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()
	router.InitRouter(userHandler, wsHandler)

	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	log.Printf("Server starting on %s", addr)
	router.Start(addr)
}
