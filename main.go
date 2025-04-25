package main

import (
	"auth-service/db"
	"auth-service/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Настройка CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Фронтенд
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
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
		protected.POST("/notes", handlers.CreateNote)
		protected.POST("/Tasks", handlers.CreateTask)
		protected.PUT("/notes/:id", handlers.UpdateNote)
		protected.GET("/notes", handlers.GetNotes)
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
