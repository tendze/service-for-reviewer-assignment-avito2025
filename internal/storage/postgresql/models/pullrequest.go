package models

import "time"

type PullRequest struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	AuthorID  uint      `gorm:"not null"`
	Author    User      `gorm:"foreignKey:AuthorID"`
	Status    string    `gorm:"not null;default:OPEN"` // OPEN, MERGED
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`
	MergedAt  *time.Time
}
