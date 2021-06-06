package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Idx             uint64 `gorm:"primaryKey;autoIncrement:true"`
	GoogleUserId    string
	AppleUserId     string
	Email           string
	CustomizedValue float64

	Deleted gorm.DeletedAt
	Created time.Time `gorm:"default:null"`
	Updated time.Time `gorm:"default:null"`
}
