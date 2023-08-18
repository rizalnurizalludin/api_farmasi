package controllers

import (
	"gudang-obat/config"
	"gudang-obat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	result := config.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed get data users",
			"date":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    users,
	})
}
