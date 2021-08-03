package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	str := "user=postgres dbname=mercado password=rockeiro,1 host=localhost sslmode=disable"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

}

func GetDatabase() *gorm.DB {

	return db
}
