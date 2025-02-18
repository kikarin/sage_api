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

	// Rute untuk portofolio_projek
	r.GET("/portofolio", controllers.GetAllPortofolioProjek)
	r.GET("/portofolio/:id", controllers.GetPortofolioProjekByID)
	r.POST("/portofolio", controllers.CreatePortofolioProjek)
	r.PUT("/portofolio/:id", controllers.UpdatePortofolioProjek)
	r.DELETE("/portofolio/:id", controllers.DeletePortofolioProjek)

	// Rute untuk blogs
	r.GET("/blogs", controllers.GetAllBlogs)
	r.GET("/blogs/:id", controllers.GetBlogByID)
	r.POST("/blogs", controllers.CreateBlog)
	r.PUT("/blogs/:id", controllers.UpdateBlog)
	r.DELETE("/blogs/:id", controllers.DeleteBlog)
}
