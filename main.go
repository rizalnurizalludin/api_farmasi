package main

import (
	"gudang-obat/config"
	"gudang-obat/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// loadEnv()
	config.InitDatabase()
	welcome := gin.Default()
	welcome.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome To API")
	})
	router := gin.Default()
	routes.InitRoute(router)
	router.Run(getPort())
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":8080"
}
