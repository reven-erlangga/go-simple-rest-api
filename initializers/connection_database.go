package initializers

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/reven-erlangga/go-simple-rest-api/helpers"
	"github.com/reven-erlangga/go-simple-rest-api/models"
	// seederBook "github.com/reven-erlangga/go-simple-rest-api/seeders/books"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		helpers.ErrorPanic(err)
	}



	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})

	

	// seederBook.DBSeed(db, 1000)

	DB = db
}