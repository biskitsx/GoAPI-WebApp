package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"www.github.com/biskitsx/go-api/webapp-sample/model"
)

var Db *gorm.DB

func ConnectDb() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("can't connected to database")
		return
	}

	fmt.Println("connected to database")

	err = Db.AutoMigrate(&model.Author{}, &model.Book{}, &model.Category{}, &model.User{})
	if err != nil {
		fmt.Println("can't migrate")
		return
	}

	fmt.Println("migration successfully")
}
