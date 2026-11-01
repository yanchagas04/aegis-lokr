package models

import (
	"gorm.io/gorm"
)

type OperationType string

const (
	OperationCreate  OperationType = "CREATE"
	OperationUpdate  OperationType = "UPDATE"
	OperationDelete  OperationType = "DELETE"
	OperationArchive OperationType = "ARCHIVE"
	OperationRestore OperationType = "RESTORE"
)

type ItemType string

const (
	ItemTypePassword ItemType = "PASSWORD"
	ITemTypeGroup    ItemType = "GROUP"
)

type Log struct {
	gorm.Model
	Operation     OperationType `gorm:"not null"`
	ItemType      ItemType      `gorm:"not null"`
	ItemID        uint          `gorm:"not null"`
	ItemName      string        `gorm:"not null"`
	AlteredFields string        `gorm:"not null"`
}
