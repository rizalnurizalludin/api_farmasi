package controllers

import (
	"gudang-obat/config"
	"gudang-obat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	var orders []models.Order

	result := config.DB.Find(&orders)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed get data orders",
			"date":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    orders,
	})
}

func GetOrderDetail(c *gin.Context) {
	id := c.Param("id")

	var orders models.Order

	if result := config.DB.First(&orders, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    orders,
	})
}

func CreateOrders(c *gin.Context) {
	var requestInput models.Order
	var med models.Med
	c.Bind(&requestInput)

	// masukkan ke database
	result := config.DB.Create(&requestInput)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed insert data",
			"data":    nil,
		})
	}

	if result := config.DB.First(&med, requestInput.MedID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		c.JSON(http.StatusConflict, gin.H{
			"status":  false,
			"message": "meds not found",
		})
		return
	}

	med.Stock = med.Stock - requestInput.Qty
	config.DB.Save(&med)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    requestInput,
	})
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var med models.Med

	var order models.Order

	if result := config.DB.First(&order, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := config.DB.First(&med, order.MedID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		c.JSON(http.StatusConflict, gin.H{
			"status":  false,
			"message": "meds not found",
		})
		return
	}

	med.Stock = med.Stock + order.Qty
	config.DB.Save(&med)

	config.DB.Delete(&order)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    order,
	})
}
