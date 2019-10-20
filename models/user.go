package models

import (
	"time"
)

// User for user management
type User struct {
	ID        string         `json:"id"`
	Email     string         `json:"email"`
	Name      string         `json:"name"`
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
	Roles     []SecurityRole `json:"roles"`
}
