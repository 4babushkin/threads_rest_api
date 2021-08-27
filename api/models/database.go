package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {

	databaseUser := os.Getenv("databaseUser")
	databasePassword := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	databaseHost := os.Getenv("databaseHost")
	databasePort := os.Getenv("databasePort")

	// Define DB connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", databaseHost, databasePort, databaseUser, databaseName, databasePassword)
	//connect to db URI
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Instance{})

	DB = db
}
