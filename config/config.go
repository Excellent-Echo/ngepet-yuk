package config

import (
	"fmt"
	"ngepet-yuk/entity"
	"os"

	// "ngepet-yuk/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	err := godotenv.Load()

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&entity.Category{})
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Courses{})
	db.AutoMigrate(&entity.Mastery{})
	db.AutoMigrate(&entity.SubType{})
	db.AutoMigrate(&entity.UserDetail{})
	db.AutoMigrate(&entity.UserTransaction{})

	return db
}
