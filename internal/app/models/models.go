package models

import "time"

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type MeritDemerit struct {
	ID     uint      `gorm:"primary_key" json:"id"`
	UserID uint      `json:"user_id"`
	Points int       `json:"points"`
	Reason string    `json:"reason"`
	Date   time.Time `json:"date"`
}
