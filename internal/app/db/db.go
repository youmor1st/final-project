package db

import (
	"final-project/internal/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "user=username dbname=mydb sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Применение миграций
	DB.AutoMigrate(&models.User{}, &models.MeritDemerit{})
}
