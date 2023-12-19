package main

import (
	"log"

	"github.com/edsonjuniordev/infra/database"
	"github.com/edsonjuniordev/internal/user"
	"github.com/edsonjuniordev/internal/ws"
	"github.com/edsonjuniordev/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database: %+v", err)
	}

	userRepo := user.NewRepository(db.GetDB())
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	go hub.Run()
	wsHandler := ws.NewHandler(hub)

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:3000")
}
