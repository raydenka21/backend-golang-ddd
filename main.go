package main

import (
	"backend-app/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
	"log"
)

var r *gin.Engine

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r = gin.Default()
	server.RegisterAPIService(r)
}

func main() {
	_ = r.Run()
}



