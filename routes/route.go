package routes

import (
	"gudang-obat/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to API")
	})
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("")
	auth.Use(controllers.JwtAuthMiddleware())

	auth.GET("/users", controllers.GetUsers)

	auth.POST("/meds", controllers.AddMeds)
	auth.GET("/meds", controllers.GetMeds)
	auth.GET("/meds/:id", controllers.GetMedDetail)
	auth.POST("/meds/:id", controllers.UpdateMed)
	auth.DELETE("meds/:id", controllers.DeleteMed)

	auth.POST("/orders", controllers.CreateOrders)
	auth.GET("/orders", controllers.GetOrders)
	auth.GET("/orders/:id", controllers.GetOrderDetail)
	auth.DELETE("orders/:id", controllers.DeleteOrder)
}
