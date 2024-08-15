// main.go
package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db = setupDB()
	defer db.Close()

	e := echo.New()

	// Основной Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS конфигурация
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	// Пути для обработчиков
	e.GET("/api/comments", getCommentsHandler)
	e.POST("/api/comments", createCommentHandler)
	e.GET("/api/comments/:id", getCommentHandler)
	e.PATCH("/api/comments/:id", updateCommentHandler)
	e.DELETE("/api/comments/:id", deleteCommentHandler)

	// Запуск сервера
	log.Fatal(e.Start(":8080"))
}
