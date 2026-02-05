package controllers

import (
	"net/http"
	"path/filepath"
	"sijaku-hebat/dtos"
	"sijaku-hebat/helpers"
	"sijaku-hebat/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	service services.ProjectService
}

func NewProjectController() *ProjectController {
	service := services.NewProjectService()
	return &ProjectController{service}
}

func (c *ProjectController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to get data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success get all data", data)
}

func (c *ProjectController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid data id", nil)
		return
	}

	data, err := c.service.GetById(uint(id))
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusNotFound, "data not found", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success get data by id", data)
}

func (c *ProjectController) Create(ctx *gin.Context) {
	var req dtos.CreateProjectDTO
	if err := ctx.ShouldBind(&req); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "image is required", err.Error())
		return
	}

	uploadDir := "uploads/project"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	filename := timestamp + "_" + file.Filename
	filePath := filepath.Join(uploadDir, filename)

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to upload image", err.Error())
		return
	}

	data, err := c.service.Create(req, filePath)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, "failed to create data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "data created successfully", data)
}

func (c *ProjectController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid data id", nil)
		return
	}

	var req dtos.UpdateProjectDTO
	if err := ctx.ShouldBind(&req); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	var imagePath string
	file, err := ctx.FormFile("image")
	if err == nil {
		uploadDir := "uploads/project"
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)

		filename := timestamp + "_" + file.Filename
		filePath := filepath.Join(uploadDir, filename)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to upload image", err.Error())
			return
		}

		imagePath = filePath
	}

	data, err := c.service.Update(req, imagePath, uint(id))
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, "failed to update data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "data updated successfully", data)
}

func (c *ProjectController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid data id", nil)
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, "failed to delete data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "data deleted successfully", nil)
}
