package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Nis      string `gorm:"unique;not null" json:"nis"`
	Password string `json:"-"`
	Class    string `json:"class"`
	Phone    string `json:"phone"`
	// Company  string `json:"company"`
	IsAdmin string `json:"is_admin"`
}
