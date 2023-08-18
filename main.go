package main

import (
	"gudang-obat/config"
	"gudang-obat/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	config.InitDatabase()
	router := gin.Default()
	routes.InitRoute(router)
	router.Run(getPort())
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8080"
}
