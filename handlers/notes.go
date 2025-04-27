package handlers

import (
	"auth-service/db"
	"auth-service/models"
	"log"
	"net/http"
	"time"

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания заметки"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Заметка создана", "note": note})
}

func UpdateNote(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неавторизован"})
		return
	}

	noteID := c.Param("id")
	var note models.Note
	if err := db.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заметка не найдена"})
		return
	}

	var input models.Note
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	note.Content = input.Content
	note.Title = input.Title
	db.DB.Save(&note)

	c.JSON(http.StatusOK, gin.H{"message": "Заметка обновлена", "note": note})
}

func GetNotes(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не автоирзован"})
		return
	}

	var notes []models.Note
	if err := db.DB.Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить заметки"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"notes": notes})
}

func DeleteNote(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}
	noteID := c.Param("id")
	var note models.Note
	if err := db.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		log.Printf("Ошибка поиска заметки: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Заметка не найдена"})
		return
	}

	// Мягкое удаление: устанавливаем deleted_at
	if err := db.DB.Model(&note).Update("deleted_at", time.Now()).Error; err != nil {
		log.Printf("Ошибка при установке deleted_at: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить заметку"})
		return
	}

	log.Printf("Заметка ID: %s успешно удалена (мягкое удаление)", noteID)
	c.JSON(http.StatusOK, gin.H{"message": "Заметка удалена"})
}
