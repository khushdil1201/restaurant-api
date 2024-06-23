package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "restaurant-api/models"
)

// @Summary Get menu
// @Description Get the list of menu items
// @Tags menu
// @Accept json
// @Produce json
// @Success 200 {array} models.MenuItem
// @Router /menu [get]
func GetMenu(c *gin.Context) {
    var menu []models.MenuItem
    models.DB.Find(&menu)
    c.JSON(http.StatusOK, menu)
}

// @Summary Add menu item
// @Description Add a new menu item
// @Tags menu
// @Accept json
// @Produce json
// @Param item body models.MenuItem true "Menu item"
// @Success 200 {object} models.MenuItem
// @Router /menu [post]
func AddMenuItem(c *gin.Context) {
    var input models.MenuItem

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Create(&input)
    c.JSON(http.StatusOK, input)
}

// @Summary Delete menu item
// @Description Delete a menu item by ID
// @Tags menu
// @Accept json
// @Produce json
// @Param id path int true "Menu item ID"
// @Success 200 {object} map[string]bool
// @Router /menu/{id} [delete]
func DeleteMenuItem(c *gin.Context) {
    var menuItem models.MenuItem
    if err := models.DB.Where("id = ?", c.Param("id")).First(&menuItem).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found!"})
        return
    }

    models.DB.Delete(&menuItem)
    c.JSON(http.StatusOK, gin.H{"data": true})
}
