package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tongzcomradez/go-gorm/db"
	"github.com/tongzcomradez/go-gorm/model"
)

func main() {
	db := connectDB()
	defer func() {
		db.Close()
	}()

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Profile{})
	fmt.Println("Migrate Sucess")
}

func connectDB() *gorm.DB {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	return db
}
