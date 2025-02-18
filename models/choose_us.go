package models

import "gorm.io/gorm"

type ChooseUs struct {
	gorm.Model
	Title     string `json:"title"`
	Deskripsi string `json:"deskripsi"`
}
