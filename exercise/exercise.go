package exercise

import "github.com/jinzhu/gorm"

// Exercise data model
type Exercise struct {
	gorm.Model
	Name        string
	Description string
}
