package database

import (
	"fmt"
	"user/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(port string, host string, user string, password string, name string) {
	fmt.Println("InitDB")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, name, port,
	)

	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database", err)
		panic("Failed to connect to database")
	}

	fmt.Println("Pre migrate")
	db.AutoMigrate(&models.User{})
	fmt.Println("Post migrate")
	DB = db
}
