// crud.go
package main

import (
	"github.com/google/uuid"
)

// Создать новый коментарий
func createComment(comment Comment) error {
	comment.ID = uuid.New()
	_, err := db.Exec(`
        INSERT INTO comments (id, author, email, telephone, text, visibility, created_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		comment.ID, comment.Author, comment.Email, comment.Telephone, comment.Text, comment.Visibility, comment.CreatedAt)
	return err
}

// Получить все коментарии
func getComments() ([]Comment, error) {
	rows, err := db.Query("SELECT id, author, email, telephone, text, visibility, created_at FROM comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.ID, &c.Author, &c.Email, &c.Telephone, &c.Text, &c.Visibility, &c.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

// Получить конкретный коментарий по id(uuid)
func getComment(id uuid.UUID) (Comment, error) {
	var c Comment
	err := db.QueryRow("SELECT id, author, email, telephone, text, visibility, created_at FROM comments WHERE id = $1", id).
		Scan(&c.ID, &c.Author, &c.Email, &c.Telephone, &c.Text, &c.Visibility, &c.CreatedAt)
	return c, err
}

// Изменить конкретный коментарий (created_at изменить нельзя)
// Update a specific comment (created_at cannot be changed)
func updateComment(id uuid.UUID, author *string, email *string, telephone *string, text *string, visibility bool) error {
	_, err := db.Exec(`
		UPDATE comments 
		SET 
			author = COALESCE($1, author),
			email = COALESCE($2, email),
			telephone = COALESCE($3, telephone),
			text = COALESCE($4, text),
			visibility = $5 
		WHERE id = $6`,
		author, email, telephone, text, visibility, id)
	return err
}

// Удалить конкретный коментарий
func deleteComment(id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM comments WHERE id = $1", id)
	return err
}

// Create a new sum
func createSum(sum Sum) error {
	sum.ID = uuid.New()
	_, err := db.Exec(`
        INSERT INTO sums (id, name, amount)
        VALUES ($1, $2, $3)`,
		sum.ID, sum.Name, sum.Amount)
	return err
}

// Get all sums
func getSums() ([]Sum, error) {
	rows, err := db.Query("SELECT id, name, amount FROM sums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sums []Sum
	for rows.Next() {
		var s Sum
		if err := rows.Scan(&s.ID, &s.Name, &s.Amount); err != nil {
			return nil, err
		}
		sums = append(sums, s)
	}
	return sums, nil
}

// Get a specific sum by id (uuid)
func getSum(id uuid.UUID) (Sum, error) {
	var s Sum
	err := db.QueryRow("SELECT id, name, amount FROM sums WHERE id = $1", id).
		Scan(&s.ID, &s.Name, &s.Amount)
	return s, err
}

// Update a specific sum
func updateSum(id uuid.UUID, name string, amount int32) error {
	_, err := db.Exec("UPDATE sums SET name = $1, amount = $2 WHERE id = $3", name, amount, id)
	return err
}

// Delete a specific sum
func deleteSum(id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM sums WHERE id = $1", id)
	return err
}
