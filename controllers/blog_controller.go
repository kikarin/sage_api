package controllers

import (
	"net/http"
	"strconv"

	"sage_api/config"
	"sage_api/models"

	"github.com/gin-gonic/gin"
)

// GetAllBlogs mengambil semua data blogs dengan pagination
func GetAllBlogs(c *gin.Context) {
	// Ambil parameter `page` dari query (default = 1)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 10 // Batasi 10 item per halaman

	// Hitung offset
	offset := (page - 1) * limit

	var blogs []models.Blog
	var total int64

	// Hitung total data
	config.DB.Model(&models.Blog{}).Count(&total)

	// Ambil data dengan limit dan offset
	result := config.DB.Limit(limit).Offset(offset).Find(&blogs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Hitung total halaman
	totalPages := (total + int64(limit) - 1) / int64(limit)

	// Format response
	response := gin.H{
		"current_page": page,
		"data":         blogs,
		"message":      "Berhasil Mendapatkan Data",
		"success":      true,
		"total":        total,
		"total_page":   totalPages,
	}

	c.JSON(http.StatusOK, response)
}

// GetBlogByID mengambil satu data berdasarkan ID
func GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	result := config.DB.First(&blog, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// CreateBlog untuk menambahkan data baru
func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.DB.Create(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Data created successfully", "data": blog})
}

// UpdateBlog untuk mengupdate data berdasarkan ID
func UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	// Cek apakah data ada
	result := config.DB.First(&blog, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	// Bind data baru
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&blog)
	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully", "data": blog})
}

// DeleteBlog untuk menghapus data berdasarkan ID
func DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	// Cek apakah data ada
	result := config.DB.First(&blog, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	config.DB.Delete(&blog)
	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
}
