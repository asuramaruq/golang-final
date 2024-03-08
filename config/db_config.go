package config

import (
	"fmt"
	"golang_article/model/entity"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDBConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT") // Assuming you have DB_PORT in your .env

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed connecting to database")
	}

	//Migrate entity model to database
	errMigration := db.AutoMigrate(&entity.Article{}, &entity.User{})
	if errMigration != nil {
		return nil
	}
	return db
}

// CloseDBConnection is a function to close database connection
func CloseDBConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}

	errCloseDB := dbSQL.Close()
	if errCloseDB != nil {
		panic("Fail to close connection")
	}
}
