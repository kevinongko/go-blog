package model

import "github.com/jinzhu/gorm"

// Article is life
type Article struct {
	Title string
	Body  string
	gorm.Model
}
