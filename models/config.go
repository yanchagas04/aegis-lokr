package models

import (
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Salt     []byte `gorm:"not null"`
	Verifier []byte `gorm:"not null"`
	Version  uint   `gorm:"not null;default:1"`
}
