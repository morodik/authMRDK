package handlers

import (
	"auth-service/db"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не зарегистрирован"})
		return
	}

	var note models.Note
	if err := c.ShouldBind(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ввод"})
		return
	}
	note.UserID = userID.(uint)
	if err := db.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания задачи"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Задача создана", "note": note})
}
