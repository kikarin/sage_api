package controllers

import (
	"net/http"
	"sage_api/config"
	"sage_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllFaqs untuk mengambil daftar FAQ dengan pagination
func GetAllFaqs(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10

	// Hitung offset
	offset := (page - 1) * limit

	var faqs []models.Faq
	var total int64

	// Hitung total data
	config.DB.Model(&models.Faq{}).Count(&total)

	// Ambil data dengan limit dan offset
	result := config.DB.Limit(limit).Offset(offset).Find(&faqs)
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
		"data":         faqs,
		"message":      "Berhasil mengambil data FAQ",
		"success":      true,
	})
}

// CreateFaq untuk menambahkan FAQ baru
func CreateFaq(c *gin.Context) {
	var faq models.Faq
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&faq)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "FAQ berhasil ditambahkan", "data": faq})
}

// GetFaqByID untuk mengambil FAQ berdasarkan ID
func GetFaqByID(c *gin.Context) {
	id := c.Param("id")
	var faq models.Faq
	result := config.DB.First(&faq, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, faq)
}

// UpdateFaq untuk memperbarui FAQ berdasarkan ID
func UpdateFaq(c *gin.Context) {
	id := c.Param("id")
	var faq models.Faq

	// Cek apakah FAQ ada
	result := config.DB.First(&faq, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ tidak ditemukan"})
		return
	}

	// Bind data baru
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&faq)
	c.JSON(http.StatusOK, gin.H{"message": "FAQ berhasil diperbarui", "data": faq})
}

// DeleteFaq untuk menghapus FAQ berdasarkan ID
func DeleteFaq(c *gin.Context) {
	id := c.Param("id")
	var faq models.Faq

	// Cek apakah FAQ ada
	result := config.DB.First(&faq, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ tidak ditemukan"})
		return
	}

	config.DB.Delete(&faq)
	c.JSON(http.StatusOK, gin.H{"message": "FAQ berhasil dihapus"})
}
