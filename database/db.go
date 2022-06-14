package database

import (
	"fmt"
	"pratice-sesi8/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "tugas-8"
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success connect to DB using GORM")
	db.Debug().AutoMigrate(&models.Items{}, &models.Orders{})
	fmt.Println("Success migrate table")
}

func GetDB() *gorm.DB {
	return db
}
