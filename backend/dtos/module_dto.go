package dtos

type CreateModuleDTO struct {
	Title string `form:"title" binding:"required"`
	Desc  string `form:"desc" binding:"required"`
	ItcID uint   `form:"itc_id" binding:"required"`
}

type UpdateModuleDTO struct {
	Title *string `form:"title"`
	Desc  *string `form:"desc"`
	ItcID *uint   `form:"itc_id"`
}
