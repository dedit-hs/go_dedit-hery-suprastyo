package config

import (
	"fmt"

	"github.com/dedit-hs/go_dedit-hery-suprastyo/10_orm_and_code_structure/Praktikum/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	config := Config{
		DB_Username: "root",
		DB_Password: "masuklah",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "jekom",
	}
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	migrateDB()
}

func migrateDB() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}
