package controllers

import (
	"gudang-obat/config"
	"gudang-obat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateMedRequestBody struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Price uint   `json:"price"`
	Stock uint   `json:"stock"`
}

func GetMeds(c *gin.Context) {
	var meds []models.Med

	result := config.DB.Find(&meds)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed get data meds",
			"date":    nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    meds,
	})
}

func GetMedDetail(c *gin.Context) {
	id := c.Param("id")

	var meds models.Med

	if result := config.DB.First(&meds, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    meds,
	})
}

func AddMeds(c *gin.Context) {
	var requestInput models.Med
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

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    requestInput,
	})
}

func UpdateMed(c *gin.Context) {
	id := c.Param("id")
	body := UpdateMedRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var med models.Med

	if result := config.DB.First(&med, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	med.Name = body.Name
	med.Type = body.Type
	med.Price = body.Price
	med.Stock = body.Stock

	config.DB.Save(&med)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    med,
	})
}

func DeleteMed(c *gin.Context) {
	id := c.Param("id")

	var med models.Med

	if result := config.DB.First(&med, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	config.DB.Delete(&med)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Berhasil",
		"data":    med,
	})
}
