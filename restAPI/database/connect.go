package database

import (
	"github.com/DucHoangManh/golabs/restAPI/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB
func Connect() {
	db ,err := gorm.Open(mysql.Open("mysqluser:1@/store_sample"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = db
	db.AutoMigrate(&models.User{}, 
		&models.Category{}, 
		&models.Product{}, 
		&models.Image{}, 
		&models.Review{}, 
		&models.Price{},
		&models.Cart{})
}