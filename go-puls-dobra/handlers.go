package main

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Handler to get all comments
func getCommentsHandler(c echo.Context) error {
	comments, err := getComments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch comments"})
	}
	return c.JSON(http.StatusOK, comments)
}

// Handler to create a new comment
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

// Handler to get a specific comment by ID
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

// Handler to update a specific comment
func updateCommentHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	var updateData struct {
		Author     *string `json:"author,omitempty"`
		Email      *string `json:"email,omitempty"`
		Telephone  *string `json:"telephone,omitempty"`
		Text       *string `json:"text,omitempty"`
		Visibility bool    `json:"visibility"`
	}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := updateComment(id, updateData.Author, updateData.Email, updateData.Telephone, updateData.Text, updateData.Visibility); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to update comment"})
	}
	return c.NoContent(http.StatusNoContent)
}

// Handler to delete a specific comment
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

// Handler to get all sums
func getSumsHandler(c echo.Context) error {
	sums, err := getSums()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch sums"})
	}
	return c.JSON(http.StatusOK, sums)
}

// Handler to create a new sum
func createSumHandler(c echo.Context) error {
	var sum Sum
	if err := c.Bind(&sum); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := createSum(sum); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to create sum"})
	}
	return c.JSON(http.StatusCreated, sum)
}

// Handler to get a specific sum by ID
func getSumHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	sum, err := getSum(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to fetch sum"})
	}
	return c.JSON(http.StatusOK, sum)
}

// Handler to update a specific sum
func updateSumHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	var updateData struct {
		Name   string `json:"name"`
		Amount int32  `json:"amount"`
	}
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
	if err := updateSum(id, updateData.Name, updateData.Amount); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to update sum"})
	}
	return c.NoContent(http.StatusNoContent)
}

// Handler to delete a specific sum
func deleteSumHandler(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	if err := deleteSum(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to delete sum"})
	}
	return c.NoContent(http.StatusNoContent)
}
