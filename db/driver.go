package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func MigrateDB(db *gorm.DB, models ...any) {
	err := db.AutoMigrate(models...)
	if err != nil {
		panic("Failed to migrate database")
	}
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database connection")
	}
	if err := sqlDB.Close(); err != nil {
		panic("Failed to close database")
	}
}
