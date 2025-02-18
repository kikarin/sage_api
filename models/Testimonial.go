package models

import (
	"gorm.io/gorm"
)

// Testimonial struct untuk representasi tabel 'testimonials' di database
type Testimonial struct {
	gorm.Model
	Deskripsi string `json:"deskripsi"`
	Rating    int    `json:"rating"`
}
