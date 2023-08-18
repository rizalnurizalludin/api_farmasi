package routes

import (
	"gudang-obat/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to API")
	})
	r.POST("/api/register", controllers.Register)
	r.POST("api//login", controllers.Login)
	r.GET("/api/meds", controllers.GetMeds)

	auth := r.Group("")
	auth.Use(controllers.JwtAuthMiddleware())

	auth.GET("/users", controllers.GetUsers)

	auth.POST("/meds", controllers.AddMeds)
	auth.GET("/meds", controllers.GetMeds)
	auth.GET("/meds/:id", controllers.GetMedDetail)
	auth.POST("/meds/:id", controllers.UpdateMed)
	auth.DELETE("meds/:id", controllers.DeleteMed)
}
