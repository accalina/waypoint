package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("waypoint.db"), &gorm.Config{})

	if err != nil {
		panic("cannot connect DB")
	}

	database.AutoMigrate(&Todo{})
	DB = database
}
