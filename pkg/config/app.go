package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect initializes a connection to the MySQL database.
func Connect() {
	// Replace USER_NAME, PASSWORD, and NAME_OF_TABLE with your actual MySQL database credentials and table name.
	d, err := gorm.Open("mysql", "USER_NAME:PASSWORD/NAME_OF_TABLE?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB returns the instance of the connected gorm.DB.
func GetDB() *gorm.DB {
	return db
}
