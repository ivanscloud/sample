package config

import (
	"alta-store/models"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var HTTP_PORT int

func InitDb() {
	connectionString := os.Getenv("CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}
func InitPort() {
	var err error
	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.Customers{})
	DB.AutoMigrate(&models.Categories{})
	DB.AutoMigrate(&models.Products{})
	DB.AutoMigrate(&models.Couriers{})
	DB.AutoMigrate(&models.Orders{})
	DB.AutoMigrate(&models.Carts{})
	DB.AutoMigrate(&models.Cartitems{})
	DB.AutoMigrate(&models.Checkout_items{})
	DB.AutoMigrate(&models.Payments{})
}
