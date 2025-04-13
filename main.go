package main

import (
	"auth-service/db"
	"auth-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Инициализация базы данных
	db.Init()

	// Регистрация маршрутов
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/logout", handlers.Logout)

	// Защищенные маршруты
	protected := router.Group("/api")
	protected.Use(handlers.AuthMiddleware())
	{
		protected.POST("/Notes", handlers.CreateNote)
	}

	// Тестовый маршрут
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "aboba",
		})
	})

	// Запуск сервера
	router.Run(":8080")
}
