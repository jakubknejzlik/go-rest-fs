package model

import "github.com/jinzhu/gorm"

// File ...
type File struct {
	gorm.Model
	Code string
	Size uint
}
