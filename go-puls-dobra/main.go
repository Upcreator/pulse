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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	// Comment routes
	e.GET("/api/comments", getCommentsHandler)
	e.POST("/api/comments", createCommentHandler)
	e.GET("/api/comments/:id", getCommentHandler)
	e.PATCH("/api/comments/:id", updateCommentHandler)
	e.DELETE("/api/comments/:id", deleteCommentHandler)

	// Sum routes
	e.GET("/api/sums", getSumsHandler)
	e.POST("/api/sums", createSumHandler)
	e.GET("/api/sums/:id", getSumHandler)
	e.PATCH("/api/sums/:id", updateSumHandler)
	e.DELETE("/api/sums/:id", deleteSumHandler)

	// Start the server
	log.Fatal(e.Start(":8080"))
}
