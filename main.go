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

	// Setup router
	r := gin.Default()
	routes.InitRoutes(r)

	fmt.Println("🚀 Server running at http://localhost:8081")

	// Jalankan server di port 8081
	r.Run(":8081")
}
