package models

import "gorm.io/gorm"

// PortofolioProjek struct untuk representasi tabel portofolio_projek di database
type PortofolioProjek struct {
	gorm.Model
	Foto           string `json:"foto"`
	NamaWeb        string `json:"nama_web"`
	NamaPerusahaan string `json:"nama_perusahaan"`
}

// TableName mengatur nama tabel agar sesuai dengan database
func (PortofolioProjek) TableName() string {
	return "portofolio_projek"
}
