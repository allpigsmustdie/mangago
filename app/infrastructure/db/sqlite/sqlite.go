package sqlite

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InMemory() (*gorm.DB, error) {
	return gorm.Open("sqlite3", ":memory:")
}