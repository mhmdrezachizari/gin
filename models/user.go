package models

import "time"

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	NationalID   string    `json:"national_id" gorm:"uniqueIndex;size:10;not null"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}