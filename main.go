package main

import (
	"go-redis-caching/pkg/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// User endpoints
	userHandler := handler.NewUserHandler()
	r.POST("/user", userHandler.AddUser)
	r.GET("/user/:id", userHandler.GetUser)

	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
