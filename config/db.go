package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
}
