// handlers.go
package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Функция обработчкик получения всех коментариев
func getCommentsHandler(c echo.Context) error {
	comments, err := getComments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch comments"})
	}
	return c.JSON(http.StatusOK, comments)
}

// Функция обработчкик получения всех коментариев
func createCommentHandler(c echo.Context) error {
	var comment Comment
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	comment.CreatedAt = time.Now() // Explicitly setting the CreatedAt field (optional)
	if err := createComment(comment); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to create comment"})
	}
	return c.JSON(http.StatusCreated, comment)
}

// Функция обработчик для получения конкретного коментария
func getCommentHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	comment, err := getComment(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch comment"})
	}
	return c.JSON(http.StatusOK, comment)
}

// Функция обработчик для конкретного коментария
func updateCommentHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	var updateData struct {
		Author     string `json:"author"`
		Text       string `json:"text"`
		Visibility bool   `json:"visibility"`
	}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := updateComment(id, updateData.Author, updateData.Text, updateData.Visibility); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to update comment"})
	}
	return c.NoContent(http.StatusNoContent)
}

// Функция обработчик для удаления команетария
func deleteCommentHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	if err := deleteComment(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to delete comment"})
	}
	return c.NoContent(http.StatusNoContent)
}
