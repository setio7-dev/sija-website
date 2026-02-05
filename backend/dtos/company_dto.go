package dtos

type CreateCompanyDTO struct {
	Name  string `form:"name" binding:"required"`
	Desc  string `form:"desc" binding:"required"`
	Link  string `form:"link" binding:"required"`
	Phone string `form:"phone" binding:"required"`
	Email string `form:"email" binding:"required,email"`
}

type UpdateCompanyDTO struct {
	Name  *string `form:"name"`
	Desc  *string `form:"desc"`
	Link  *string `form:"link"`
	Phone *string `form:"phone"`
	Email *string `form:"email" binding:"omitempty,email"`
}
