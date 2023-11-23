package database

import (
	// "fmt"
	"fmt"
	"log"
	"os"

	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")

	dsn := fmt.Sprintf(`
		host=%s
		user=%s
		password=%s
		dbname=%s
		port=%s
		sslmode=disable`, dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}
	return db
}
