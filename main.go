package main

import (
	"auth-service/db"
	"auth-service/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	// инициализируем базу данных
	db.Init() // прямо вызываем Init() без проверки ошибки

	// создаем новый роутер
	router := gin.Default()

	// роуты для авторизации
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.GET("/logout", handlers.Logout)

	// роут для проверки API
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API работает",
		})
	})

	router.Run(":8080")
}
