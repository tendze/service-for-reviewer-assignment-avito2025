package models

import "time"

type PRReviewer struct {
	ID         uint      `gorm:"primaryKey"`
	PRID       uint      `gorm:"not null"`
	ReviewerID uint      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"default:now()"`
}

func (PRReviewer) TableName() string {
	return "pr_reviewer"
}
