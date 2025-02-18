package controllers

import (
	"net/http"
	"sage_api/config"
	"sage_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllChooseUs mengambil semua data choose_us dengan pagination
func GetAllChooseUs(c *gin.Context) {
	// Ambil parameter `page` dari query (default = 1)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10 // Batasi 10 item per halaman

	// Hitung offset
	offset := (page - 1) * limit

	var chooseUs []models.ChooseUs
	var total int64

	// Hitung total data
	config.DB.Model(&models.ChooseUs{}).Count(&total)

	// Ambil data dengan limit dan offset
	result := config.DB.Limit(limit).Offset(offset).Find(&chooseUs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Hitung total halaman
	totalPages := (total + int64(limit) - 1) / int64(limit)

	// Format response
	response := gin.H{
		"current_page": page,
		"total_data":   total,
		"total_page":   totalPages,
		"data":         chooseUs,
		"message":      "Berhasil Mendapatkan Data",
		"success":      true,
	}

	c.JSON(http.StatusOK, response)
}

// GetChooseUsByID mengambil satu choose_us berdasarkan ID
func GetChooseUsByID(c *gin.Context) {
	id := c.Param("id")
	var chooseUs models.ChooseUs
	result := config.DB.First(&chooseUs, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ChooseUs not found"})
		return
	}
	c.JSON(http.StatusOK, chooseUs)
}

// CreateChooseUs untuk menambahkan choose_us baru
func CreateChooseUs(c *gin.Context) {
	var chooseUs models.ChooseUs
	if err := c.ShouldBindJSON(&chooseUs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&chooseUs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "ChooseUs created successfully", "data": chooseUs})
}

// UpdateChooseUs untuk mengupdate choose_us berdasarkan ID
func UpdateChooseUs(c *gin.Context) {
	id := c.Param("id")
	var chooseUs models.ChooseUs

	// Cek apakah ChooseUs ada
	result := config.DB.First(&chooseUs, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ChooseUs not found"})
		return
	}

	// Bind data baru
	if err := c.ShouldBindJSON(&chooseUs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&chooseUs)
	c.JSON(http.StatusOK, gin.H{"message": "ChooseUs updated successfully", "data": chooseUs})
}

// DeleteChooseUs untuk menghapus choose_us berdasarkan ID
func DeleteChooseUs(c *gin.Context) {
	id := c.Param("id")
	var chooseUs models.ChooseUs

	// Cek apakah ChooseUs ada
	result := config.DB.First(&chooseUs, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ChooseUs not found"})
		return
	}

	config.DB.Delete(&chooseUs)
	c.JSON(http.StatusOK, gin.H{"message": "ChooseUs deleted successfully"})
}
