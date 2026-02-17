package main

import (
	"log"

	"github.com/evamdf/api-project/database"
	"github.com/evamdf/api-project/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDatabase()

	r := gin.Default()
	routes.Setup(r)
	r.Run(":8080")
}
