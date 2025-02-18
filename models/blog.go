package models

import "gorm.io/gorm"

// Blog struct untuk representasi tabel blogs di database
type Blog struct {
	gorm.Model
	Foto      string `json:"foto"`
	Title     string `json:"title"`
	Deskripsi string `json:"deskripsi"`
}

// TableName mengatur nama tabel agar sesuai dengan database
func (Blog) TableName() string {
	return "blogs"
}
