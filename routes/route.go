package routes

import (
	"gudang-obat/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	public := r.Group("/")

	public.GET("/", func(c *gin.Context) {
		c.JSON(200, "Welcome to API")
	})
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(controllers.JwtAuthMiddleware())

	protected.GET("/users", controllers.GetUsers)

	protected.POST("/meds", controllers.AddMeds)
	protected.GET("/meds", controllers.GetMeds)
	protected.GET("/meds/:id", controllers.GetMedDetail)
	protected.POST("/meds/:id", controllers.UpdateMed)
	protected.DELETE("meds/:id", controllers.DeleteMed)
}
