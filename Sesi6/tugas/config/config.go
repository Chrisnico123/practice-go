package config

import (
	"fmt"
	"movies/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_USER = "DB_USER"
	DB_PASS = "DB_PASS"
	DB_NAME = "DB_NAME"
)

func GetString(key string) string {
	return os.Getenv(key)
}

func ConnectDb() (*gorm.DB , error) {
	var (
		DB_Host   = GetString(DB_HOST)
		DB_Port   = GetString(DB_PORT)
		DB_User   = GetString(DB_USER)
		DB_Pass   = GetString(DB_PASS)
		DB_DbName = GetString(DB_NAME)
	)

	fmt.Println(DB_Host)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_Host, DB_Port, DB_User, DB_Pass, DB_DbName,
	)
	db , err := gorm.Open(postgres.Open(dsn) , &gorm.Config{})

	if err!= nil {
		return nil , err
	}

	db.Debug().AutoMigrate(models.Movies{})

	return db , nil
}