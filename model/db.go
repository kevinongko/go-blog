package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

// DB is the connection handle for the database
var DB *gorm.DB

// InitDB open connection to the database
func InitDB(dataSource string) {
	var err error
	DB, err = gorm.Open("mysql", dataSource)
	if err != nil {
		log.Panic(err)
	}
}
