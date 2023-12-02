package database

import (
	"fmt"
	"os"

	"github.com/andresroto/go-login/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// Setup database connection
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbname)
	if dsn == "" {
		panic("One or more DB environment variables are not set")
	}

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	// Migrate the database schema
	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic("Could not migrate the database")
	}
}
