package models

import "time"

type TeamCard struct {
	ID            int       `json:"id" db:"id"`
	Foto          string    `json:"foto" db:"foto"`
	Nama          string    `json:"nama" db:"nama"`
	Profesi       string    `json:"profesi" db:"profesi"`
	LinkFacebook  *string   `json:"link_facebook,omitempty" db:"link_facebook"`
	LinkTwitter   *string   `json:"link_twitter,omitempty" db:"link_twitter"`
	LinkLinkedin  *string   `json:"link_linkedin,omitempty" db:"link_linkedin"`
	LinkInstagram *string   `json:"link_instagram,omitempty" db:"link_instagram"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
