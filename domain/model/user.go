package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Idx             uint `gorm:"primaryKey"`
	GoogleUserId    string
	AppleUserId     string
	Email           string
	CustomizedValue float64

	deleted gorm.DeletedAt
	created time.Time
	updated time.Time
}
