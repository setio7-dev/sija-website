package models

import "time"

type Project struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Image     string    `json:"image"`
	Link      string    `json:"link"`
	Itc       []Itc     `gorm:"foreignKey:ProjectID" json:"itc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
