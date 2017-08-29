package database

import (
	"fmt"
	"github.com/dpitkevics/GamingPlatform/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var db *gorm.DB

func InitDatabase() {
	connection, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=gaming_platform sslmode=disable password=postgres")
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %#v", err.Error()))
	}

	db = connection

}

func GetDatabase() *gorm.DB {
	if db == nil {
		InitDatabase()
	}

	return db
}

func CloseDatabase() {
	if db != nil {
		db.Close()
	}
}

func SeedDatabase() {
	db := GetDatabase()

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Sport{})

	db.Create(&models.Sport{
		Name: "Table Tennis",
		PlayerCountInTeam: 1,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
	db.Create(&models.Sport{
		Name: "Table Football 1v1",
		PlayerCountInTeam: 1,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
	db.Create(&models.Sport{
		Name: "Table Football 2v2",
		PlayerCountInTeam: 2,
		Model: models.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
}
