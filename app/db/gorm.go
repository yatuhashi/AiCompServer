package db

import (
	// _ "github.com/go-sql-driver/mysql"
	"AiCompServer/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
	"log"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("sqlite3", dbInfoString())
	if err != nil {
		log.Panicf("Failed to connect to database: %v\n", err)
	}

	db.DB()
	db.AutoMigrate(&models.User{})
	DB = db
	// DB.Create(&models.User{Username: "tester"})
}

func dbInfoString() string {
	s, b := revel.Config.String("db.info")
	if !b {
		log.Panicf("database info not found")
	}
	return s
}
