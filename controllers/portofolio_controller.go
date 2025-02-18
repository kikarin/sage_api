package controllers

import (
	"net/http"
	"strconv"

	"sage_api/config"
	"sage_api/models"

	"github.com/gin-gonic/gin"
)

// GetAllPortofolioProjek mengambil semua data dengan pagination
func GetAllPortofolioProjek(c *gin.Context) {
	// Ambil parameter `page` dari query (default = 1)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10 // Batasi 10 item per halaman

	// Hitung offset
	offset := (page - 1) * limit

	var portofolio []models.PortofolioProjek
	var total int64

	// Hitung total data
	config.DB.Model(&models.PortofolioProjek{}).Count(&total)

	// Ambil data dengan limit dan offset
	result := config.DB.Limit(limit).Offset(offset).Find(&portofolio)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Hitung total halaman
	totalPages := (total + int64(limit) - 1) / int64(limit)

	// Format response
	response := gin.H{
		"current_page": page,
		"data":         portofolio,
		"message":      "Berhasil Mendapatkan Data",
		"success":      true,
		"total":        total,
		"total_page":   totalPages,
	}

	c.JSON(http.StatusOK, response)
}

// GetPortofolioProjekByID mengambil satu data berdasarkan ID
func GetPortofolioProjekByID(c *gin.Context) {
	id := c.Param("id")
	var portofolio models.PortofolioProjek
	result := config.DB.First(&portofolio, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}
	c.JSON(http.StatusOK, portofolio)
}

// CreatePortofolioProjek untuk menambahkan data baru
func CreatePortofolioProjek(c *gin.Context) {
	var portofolio models.PortofolioProjek
	if err := c.ShouldBindJSON(&portofolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&portofolio)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Data created successfully", "data": portofolio})
}

// UpdatePortofolioProjek untuk mengupdate data berdasarkan ID
func UpdatePortofolioProjek(c *gin.Context) {
	id := c.Param("id")
	var portofolio models.PortofolioProjek

	// Cek apakah data ada
	result := config.DB.First(&portofolio, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	// Bind data baru
	if err := c.ShouldBindJSON(&portofolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&portofolio)
	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully", "data": portofolio})
}

// DeletePortofolioProjek untuk menghapus data berdasarkan ID
func DeletePortofolioProjek(c *gin.Context) {
	id := c.Param("id")
	var portofolio models.PortofolioProjek

	// Cek apakah data ada
	result := config.DB.First(&portofolio, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	config.DB.Delete(&portofolio)
	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
}
