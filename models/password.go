package models

import (
	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	Uuid        uint   `gorm:"unique;not null"`
	Title       string `gorm:"not null"`
	Password    []byte `gorm:"not null"`
	Description *string
	GroupID     *uint
	Archived    bool `gorm:"default:false"`
}
