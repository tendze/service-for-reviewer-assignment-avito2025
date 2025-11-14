package domain

import "time"

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	TeamID    uint      `json:"team_id"`
	CreatedAt time.Time `json:"created_at"`
}
