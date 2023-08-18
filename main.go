package main

import (
	"gudang-obat/config"
	"gudang-obat/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	config.InitDatabase()
	router := gin.Default()
	routes.InitRoute(router)
	router.Run(":1234")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
