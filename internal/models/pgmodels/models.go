package pgmodels

import "time"

type Team struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"default:now()"`
}

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	IsActive  bool      `gorm:"not null;default:true"`
	TeamID    uint      `gorm:"not null"`
	Team      Team      `gorm:"foreignKey:TeamID"`
	CreatedAt time.Time `gorm:"default:now()"`
}

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

type PRReviewer struct {
	ID         uint      `gorm:"primaryKey"`
	PRID       uint      `gorm:"not null"`
	ReviewerID uint      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"default:now()"`
}
