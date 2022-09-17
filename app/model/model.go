package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Goly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Goly     string `json:"goly" gorm:"unique; not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	enverr := godotenv.Load(".env")

	if enverr != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables
	dbUsername := "user=" + os.Getenv("DB_USERNAME") + " "
	dbPassword := "password=" + os.Getenv("DB_PASSWORD") + " "
	dbName := "dbname=" + os.Getenv("DB_NAME") + " "
	dbHost := "host=" + os.Getenv("DB_HOST") + " "

	dsn := dbHost + dbUsername + dbName + dbPassword + " port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}
}
