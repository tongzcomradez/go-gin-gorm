package db

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectDB abc
func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(
		"mysql",
		"root@/test?charset=utf8&parseTime=True&loc=Local",
	)

	// defer db.Close()
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	return db, nil
}
