package main

import (
	"example.com/rest-go-api/db"
	"example.com/rest-go-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegistetRoutes(server)

	server.Run(":8080")
}
