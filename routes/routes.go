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

	// Rute untuk team_cards
	r.GET("/team_cards", controllers.GetTeamCards)
	r.GET("/team_cards/:id", controllers.GetTeamCardByID)
	r.POST("/team_cards", controllers.CreateTeamCard)
	r.PUT("/team_cards/:id", controllers.UpdateTeamCard)
	r.DELETE("/team_cards/:id", controllers.DeleteTeamCard)

	// Rute untuk choose_us dengan pagination
	r.GET("/choose_us", controllers.GetAllChooseUs)
	r.GET("/choose_us/:id", controllers.GetChooseUsByID)
	r.POST("/choose_us", controllers.CreateChooseUs)
	r.PUT("/choose_us/:id", controllers.UpdateChooseUs)
	r.DELETE("/choose_us/:id", controllers.DeleteChooseUs)

	// Rute untuk FAQ
	r.GET("/faqs", controllers.GetAllFaqs)
	r.POST("/faqs", controllers.CreateFaq)
	r.GET("/faqs/:id", controllers.GetFaqByID)
	r.PUT("/faqs/:id", controllers.UpdateFaq)
	r.DELETE("/faqs/:id", controllers.DeleteFaq)

	// Rute untuk testimonial
	r.GET("/testimonials", controllers.GetAllTestimonials)
	r.POST("/testimonials", controllers.CreateTestimonial)
	r.GET("/testimonials/:id", controllers.GetTestimonialByID)
	r.PUT("/testimonials/:id", controllers.UpdateTestimonial)
	r.DELETE("/testimonials/:id", controllers.DeleteTestimonial)

}
