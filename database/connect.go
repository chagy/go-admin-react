package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"../models"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:@/go_admin_react"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{})

}
