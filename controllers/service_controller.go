package controllers

import (
	"net/http"
	"strconv"

	"sage_api/config"
	"sage_api/models"

	"github.com/gin-gonic/gin"
)

// GetAllServices mengambil semua data services dengan pagination
func GetAllServices(c *gin.Context) {
	// Ambil parameter `page` dari query (default = 1)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10 // Batasi 10 item per halaman

	// Hitung offset
	offset := (page - 1) * limit

	var services []models.Service
	var total int64

	// Hitung total data
	config.DB.Model(&models.Service{}).Count(&total)

	// Ambil data dengan limit dan offset
	result := config.DB.Limit(limit).Offset(offset).Find(&services)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Hitung total halaman
	totalPages := (total + int64(limit) - 1) / int64(limit)

	// Format response
	response := gin.H{
		"current_page": page,
		"data":         services,
		"message":      "Berhasil Mendapatkan Data",
		"success":      true,
		"total":        total,
		"total_page":   totalPages,
	}

	c.JSON(http.StatusOK, response)
}

// GetServiceByID mengambil satu service berdasarkan ID
func GetServiceByID(c *gin.Context) {
	id := c.Param("id")
	var service models.Service
	result := config.DB.First(&service, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}
	c.JSON(http.StatusOK, service)
}

// CreateService untuk menambahkan service baru
func CreateService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&service)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Service created successfully", "data": service})
}

// UpdateService untuk mengupdate service berdasarkan ID
func UpdateService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service

	// Cek apakah service ada
	result := config.DB.First(&service, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	// Bind data baru
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&service)
	c.JSON(http.StatusOK, gin.H{"message": "Service updated successfully", "data": service})
}

// DeleteService untuk menghapus service berdasarkan ID
func DeleteService(c *gin.Context) {
	id := c.Param("id")
	var service models.Service

	// Cek apakah service ada
	result := config.DB.First(&service, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Service not found"})
		return
	}

	config.DB.Delete(&service)
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}
