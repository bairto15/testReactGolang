package main

import (
	"compas/server"
	"compas/handler"
	"log"
)

func main() {
	handlers := handler.New()

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error running http Server: %s", err.Error())
	}
}