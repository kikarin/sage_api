package models

import "gorm.io/gorm"

// User struct untuk representasi tabel users di database
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}
