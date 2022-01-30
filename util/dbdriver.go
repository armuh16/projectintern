package config

import (
	"log"
	"os"
	"sejutacita/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConfig() map[string]string {
	conf, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return conf
}

var DB *gorm.DB

// Initial Database
func InitDB() {
	connect := os.Getenv("MYSQL_CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitalMigration()
}

func InitalMigration() {
	DB.AutoMigrate(&models.User{})
}
