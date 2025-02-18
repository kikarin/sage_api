package controllers

import (
	"net/http"
	"sage_api/config"
	"sage_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllTestimonials mengambil daftar testimonial dengan pagination
func GetAllTestimonials(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10

	// Hitung offset
	offset := (page - 1) * limit

	var testimonials []models.Testimonial
	var total int64

	// Hitung total data
	config.DB.Model(&models.Testimonial{}).Count(&total)

	// Ambil data dengan limit dan offset
	result := config.DB.Limit(limit).Offset(offset).Find(&testimonials)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Hitung total halaman
	totalPages := (total + int64(limit) - 1) / int64(limit)

	// Format response dengan pagination
	c.JSON(http.StatusOK, gin.H{
		"current_page": page,
		"total_page":   totalPages,
		"total_data":   total,
		"data":         testimonials,
		"message":      "Berhasil mengambil data testimonial",
		"success":      true,
	})
}

// CreateTestimonial untuk menambahkan testimonial baru
func CreateTestimonial(c *gin.Context) {
	var testimonial models.Testimonial
	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&testimonial)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Testimonial berhasil ditambahkan", "data": testimonial})
}

// GetTestimonialByID untuk mengambil testimonial berdasarkan ID
func GetTestimonialByID(c *gin.Context) {
	id := c.Param("id")
	var testimonial models.Testimonial
	result := config.DB.First(&testimonial, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, testimonial)
}

// UpdateTestimonial untuk memperbarui testimonial berdasarkan ID
func UpdateTestimonial(c *gin.Context) {
	id := c.Param("id")
	var testimonial models.Testimonial

	// Cek apakah testimonial ada
	result := config.DB.First(&testimonial, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial tidak ditemukan"})
		return
	}

	// Bind data baru
	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&testimonial)
	c.JSON(http.StatusOK, gin.H{"message": "Testimonial berhasil diperbarui", "data": testimonial})
}

// DeleteTestimonial untuk menghapus testimonial berdasarkan ID
func DeleteTestimonial(c *gin.Context) {
	id := c.Param("id")
	var testimonial models.Testimonial

	// Cek apakah testimonial ada
	result := config.DB.First(&testimonial, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimonial tidak ditemukan"})
		return
	}

	config.DB.Delete(&testimonial)
	c.JSON(http.StatusOK, gin.H{"message": "Testimonial berhasil dihapus"})
}
