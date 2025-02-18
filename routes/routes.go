package routes

import (
	"sage_api/controllers"

	"github.com/gin-gonic/gin"
)

// InitRoutes mengatur semua rute dalam aplikasi
func InitRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Sage API!"})
	})

	// Rute untuk users
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)

	// Rute untuk services
	r.GET("/services", controllers.GetAllServices)
	r.GET("/services/:id", controllers.GetServiceByID)
	r.POST("/services", controllers.CreateService)
	r.PUT("/services/:id", controllers.UpdateService)
	r.DELETE("/services/:id", controllers.DeleteService)
}
