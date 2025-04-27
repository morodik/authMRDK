package handlers

import (
	"auth-service/db"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser возвращает данные пользователя по user_id из токена
func GetUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}

	var user models.User
	if err := db.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": user.Email,
	})
}
