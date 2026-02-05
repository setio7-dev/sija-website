package dtos

type CreateUserDTO struct {
	Name       string `json:"name" binding:"required"`
	Nis        string `json:"nis" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Class      string `json:"class" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	CompanyID  *uint  `json:"company_id"`
	CategoryID *uint  `json:"category_id"`
	Status     string `json:"status" binding:"required"`
	IsAdmin    bool   `json:"is_admin"`
}

type UpdateUserDTO struct {
	Name       *string `json:"name"`
	Password   *string `json:"password"`
	Class      *string `json:"class"`
	Phone      *string `json:"phone"`
	CompanyID  *uint   `json:"company_id"`
	CategoryID *uint   `json:"category_id"`
	Status     *string `json:"status"`
	IsAdmin    *bool   `json:"is_admin"`
}

type LoginUserDTO struct {
	Nis      string `json:"nis" binding:"required"`
	Password string `json:"password" binding:"required"`
}
