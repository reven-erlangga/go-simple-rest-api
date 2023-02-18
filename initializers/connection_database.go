package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/reven-erlangga/go-simple-rest-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})

	DB = db
}