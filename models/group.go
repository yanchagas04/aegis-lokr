package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Uuid        uint   `gorm:"unique;not null"`
	Name        string `gorm:"not null"`
	Description *string
	Image       []byte
	ParentID    *uint
	Archived    bool `gorm:"default:false"`
}
