package main

import "github.com/jinzhu/gorm"

// NewDB ...
func NewDB(url string) (*gorm.DB, error) {
	return gorm.Open("sqlite3", "test.db")
}
