// crud.go
package main

import (
	"github.com/google/uuid"
)

// Создать новый коментарий
func createComment(comment Comment) error {
	comment.ID = uuid.New()
	_, err := db.Exec(`
        INSERT INTO comments (id, author, text, visibility, created_at)
        VALUES ($1, $2, $3, $4, $5)`, comment.ID, comment.Author, comment.Text, comment.Visibility, comment.CreatedAt)
	return err
}

// Получить все коментарии
func getComments() ([]Comment, error) {
	rows, err := db.Query("SELECT id, author, text, visibility, created_at FROM comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.ID, &c.Author, &c.Text, &c.Visibility, &c.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

// Получить конкретный коментарий по id(uuid)
func getComment(id uuid.UUID) (Comment, error) {
	var c Comment
	err := db.QueryRow("SELECT id, author, text, visibility, created_at FROM comments WHERE id = $1", id).Scan(&c.ID, &c.Author, &c.Text, &c.Visibility, &c.CreatedAt)
	return c, err
}

// Изменить конкретный коментарий (created_at изменить нельзя)
func updateComment(id uuid.UUID, author, text string, visibility bool) error {
	_, err := db.Exec("UPDATE comments SET author = $1, text = $2, visibility = $3 WHERE id = $4", author, text, visibility, id)
	return err
}

// Удалить конкретный коментарий
func deleteComment(id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM comments WHERE id = $1", id)
	return err
}
