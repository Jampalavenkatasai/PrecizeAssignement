// config/database.go
package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"scoreboard-go/server/models"
)

var DB *gorm.DB

func SetupDatabase() {

	dsn := "user=postgres password=Venkat@996361 dbname=cruddb host=localhost port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	DB = db
	DB.AutoMigrate(&models.ScoreCard{})
}
