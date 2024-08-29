// model.go
package main

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID         uuid.UUID `json:"id"`
	Author     string    `json:"author"`
	Email      *string   `json:"email,omitempty"`
	Telephone  *string   `json:"telephone,omitempty"`
	Text       string    `json:"text"`
	Visibility bool      `json:"visibility"`
	CreatedAt  time.Time `json:"created_at"`
}

type Sum struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Amount int32     `json:"amount"`
}

// model
type TransitionApplication struct {
	ID   uuid.UUID `json:id`
	User string    `json:user`
	Time time.Time `json:time`
}
