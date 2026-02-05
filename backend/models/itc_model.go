package models

import "time"

type Itc struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Image string `json:"image"`

	ProjectID uint     `json:"project_id"`
	Project   *Project `json:"project"`

	Module []Module `gorm:"foreignKey:ItcID" json:"module"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
