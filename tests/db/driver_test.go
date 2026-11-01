package db_test

import (
	"path/filepath"
	"testing"

	"github.com/yanchagas04/aegis-lokr/db"
	"github.com/yanchagas04/aegis-lokr/tests/utils"

	"gorm.io/gorm"
)

func openTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	var database *gorm.DB
	utils.AssertNoPanic(t, func() {
		database = db.GetDB(filepath.Join(t.TempDir(), "test.db"))
	})
	if database == nil {
		t.Fatal("Database connection failed to be established")
	}
	t.Cleanup(func() {
		utils.AssertNoPanic(t, func() {
			db.CloseDB(database)
		})
	})
	return database
}

func TestGetDB(t *testing.T) {
	openTestDB(t)
}

func TestMigrateDB(t *testing.T) {
	type testModel struct {
		gorm.Model
		Name string `gorm:"size:255"`
	}

	t.Run("AutoMigrate models", func(t *testing.T) {
		database := openTestDB(t)
		utils.AssertNoPanic(t, func() {
			db.MigrateDB(database, &testModel{Name: "TestModel"})
		})
	})
}
