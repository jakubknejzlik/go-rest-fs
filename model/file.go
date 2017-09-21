package model

import "github.com/jinzhu/gorm"

// File ...
type File struct {
	gorm.Model
	Name string
	Size uint
}
