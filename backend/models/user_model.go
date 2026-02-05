package models

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Nis      string `gorm:"unique;not null" json:"nis"`
	Password string `json:"-"`
	Class    string `json:"class"`
	Phone    string `json:"phone"`

	CompanyID *uint    `json:"company_id"`
	Company   *Company `json:"company"`

	CategoryID *uint `json:"category_id"`
	Category   *Itc  `json:"category"`

	Status  string `json:"status"`
	IsAdmin bool   `json:"is_admin"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
