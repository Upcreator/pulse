// model.go
package main

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID         uuid.UUID `json:"id"`
	Author     string    `json:"author"`
	Text       string    `json:"text"`
	Visibility bool      `json:"visibility"`
	CreatedAt  time.Time `json:"created_at"`
}
