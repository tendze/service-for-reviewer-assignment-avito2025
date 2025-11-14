package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	IsActive  bool      `gorm:"not null;default:true"`
	TeamID    uint      `gorm:"not null"`
	Team      Team      `gorm:"foreignKey:TeamID"`
	CreatedAt time.Time `gorm:"default:now()"`
}

func (User) TableName() string {
    return "user"
}