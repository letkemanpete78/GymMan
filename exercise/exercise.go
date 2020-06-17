package exercise

import "github.com/jinzhu/gorm"

// Exercise data model
type Exercise struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50);unique_index"`
	Description string `gorm:"type:varchar(50)"`
	UUID        string `gorm:"type:varchar(50);primary_key"`
}
