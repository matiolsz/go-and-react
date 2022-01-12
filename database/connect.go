package database

import (
	"github.com/matiolsz/go-and-react/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:my-secret-pw@/go_admin"), &gorm.Config{})
	if err != nil {
		panic("cant connect to db")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{})
}
