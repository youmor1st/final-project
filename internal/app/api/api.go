package api

import (
	"github.com/gin-gonic/gin"
	"github.com/youmor1st/final-project/internal/app/db"
	"github.com/youmor1st/final-project/internal/app/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Подключение к базе данных
	db.InitDB()

	// Роуты
	r.GET("/login", handlers.LoginHandler)
	r.GET("/profile", handlers.ProfileHandler)
	r.GET("/top_students", handlers.TopStudentsHandler)
	r.GET("/top_teachers", handlers.TopTeachersHandler)
	r.GET("/blog", handlers.BlogHandler)
	r.GET("/news", handlers.NewsHandler)

	return r
}
