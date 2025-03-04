package main

import (
	"fmt"
	"sage_api/config"
	"sage_api/models"
	"sage_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi database
	config.InitDB()

	// Auto-migrate tabel services
	config.DB.AutoMigrate(&models.Service{})
	config.DB.AutoMigrate(&models.PortofolioProjek{})
	config.DB.AutoMigrate(&models.Blog{})
	config.DB.AutoMigrate(&models.Blog{})
	config.DB.AutoMigrate(&models.ChooseUs{})
	config.DB.AutoMigrate(&models.Faq{})
	config.DB.AutoMigrate(&models.Testimonial{})

	// Setup router
	r := gin.Default()
	routes.InitRoutes(r)

	fmt.Println("🚀 Server running at http://localhost:8081")

	// Jalankan server di port 8081
	r.Run(":8081")
}
