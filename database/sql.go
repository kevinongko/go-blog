package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

// ORM is the connection handle for the database
var ORM *gorm.DB

// Init open connection to the database
func Init(dataSource string) {
	var err error
	ORM, err = gorm.Open("mysql", dataSource)
	if err != nil {
		log.Panic(err)
	}
}
