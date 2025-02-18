package models

import (
	"gorm.io/gorm"
)

// Faq struct untuk representasi tabel 'faqs' di database
type Faq struct {
	gorm.Model
	Pertanyaan string `json:"pertanyaan"`
	Jawaban    string `json:"jawaban"`
}
