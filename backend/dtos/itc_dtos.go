package dtos

type CreateItcDTO struct {
	Name      string `form:"name" binding:"required"`
	Desc      string `form:"desc" binding:"required"`
	ProjectID uint   `form:"project_id" binding:"required"`
}

type UpdateItcDTO struct {
	Name      *string `form:"name"`
	Desc      *string `form:"desc"`
	ProjectID *uint   `form:"project_id"`
}
