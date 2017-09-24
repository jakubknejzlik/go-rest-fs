package model

import "github.com/jinzhu/gorm"

// File ...
type File struct {
	gorm.Model
	Name string
	Size uint
}

// CreateFileInDB create row in table with filename and size
func CreateFileInDB(db *gorm.DB, name string, size uint) error {
	return db.Create(&File{Name: name, Size: size}).Error
}
