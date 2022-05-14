package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=todo_list_app password=password sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Task{})

	DB = db
}
