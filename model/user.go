package model

import "github.com/jinzhu/gorm"

// User is life
type User struct {
	Name string
	Age  string
	gorm.Model
}
