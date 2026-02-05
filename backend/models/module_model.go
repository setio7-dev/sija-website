package models

import "time"

type Module struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Itc       *Itc      `json:"itc"`
	ItcID     uint      `json:"itc_id"`
	Image     string    `json:"image"`
	File      string    `json:"file"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
