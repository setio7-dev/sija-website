package dtos

type CreateProjectDTO struct {
	Name string `form:"name" binding:"required"`
	Desc string `form:"desc" binding:"required"`
	Link string `form:"link" binding:"required"`
}

type UpdateProjectDTO struct {
	Name *string `form:"name"`
	Desc *string `form:"desc"`
	Link *string `form:"link"`
}
