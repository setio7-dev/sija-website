package models

import "time"

type Company struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
	Link  string `json:"link"`
	Phone string `json:"phone"`
	Email string `json:"email"`

	User []User `json:"users,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
