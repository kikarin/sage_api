package controllers

import (
	"net/http"
	"strconv"

	"sage_api/config"
	"sage_api/models"

	"github.com/gin-gonic/gin"
)

// GET /team_cards -> Ambil semua data (Tanpa Pagination)
func GetTeamCards(c *gin.Context) {
	var teamCards []models.TeamCard
	if err := config.DB.Order("id DESC").Find(&teamCards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "data": []models.TeamCard{}, "message": "Failed to fetch team cards", "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": teamCards, "message": "Success", "success": true})
}

// GET /team_cards/:id -> Ambil satu data berdasarkan ID
func GetTeamCardByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "data": nil, "message": "Invalid team card ID", "success": false})
		return
	}

	var teamCard models.TeamCard
	if err := config.DB.First(&teamCard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "message": "Team card not found", "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": teamCard, "message": "Success", "success": true})
}

// POST /team_cards -> Tambah data baru
func CreateTeamCard(c *gin.Context) {
	var newCard models.TeamCard
	if err := c.ShouldBindJSON(&newCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "data": nil, "message": "Invalid input data", "success": false})
		return
	}

	if err := config.DB.Create(&newCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "data": nil, "message": "Failed to create team card", "success": false})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": 201, "data": newCard, "message": "Team card created successfully", "success": true})
}

// PUT /team_cards/:id -> Update data berdasarkan ID
func UpdateTeamCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "data": nil, "message": "Invalid team card ID", "success": false})
		return
	}

	var updatedCard models.TeamCard
	if err := c.ShouldBindJSON(&updatedCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "data": nil, "message": "Invalid input data", "success": false})
		return
	}

	var teamCard models.TeamCard
	if err := config.DB.First(&teamCard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "message": "Team card not found", "success": false})
		return
	}

	if err := config.DB.Model(&teamCard).Updates(updatedCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "data": nil, "message": "Failed to update team card", "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": updatedCard, "message": "Team card updated", "success": true})
}

// DELETE /team_cards/:id -> Hapus data berdasarkan ID
func DeleteTeamCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "data": nil, "message": "Invalid team card ID", "success": false})
		return
	}

	var teamCard models.TeamCard
	if err := config.DB.First(&teamCard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "data": nil, "message": "Team card not found", "success": false})
		return
	}

	if err := config.DB.Delete(&teamCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "data": nil, "message": "Failed to delete team card", "success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "data": nil, "message": "Team card deleted", "success": true})
}
