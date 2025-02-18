package models

import "gorm.io/gorm"

// Service struct untuk representasi tabel services di database
type Service struct {
	gorm.Model
	Logo       string `json:"logo"`
	Title      string `json:"title"`
	Deskripsi  string `json:"deskripsi"`
}
