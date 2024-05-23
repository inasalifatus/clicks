package config

import (
	"clicks/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "toor1234",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "clicks",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	// connectionString := os.Getenv("CONNECTION_STRING")
	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		log.Fatal("error connecting to database:", e)
	}
	// DB.Debug().AutoMigrate(models.User{}, models.Photos{})
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Profile{})

}

func GetDB() *gorm.DB {
	return DB
}
