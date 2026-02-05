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

type CompanyController struct {
	service services.CompanyService
}

func NewCompanyController() *CompanyController {
	service := services.NewCompanyService()
	return &CompanyController{service}
}

func (c *CompanyController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to get data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success get all data", data)
}

func (c *CompanyController) GetById(ctx *gin.Context) {
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

func (c *CompanyController) Create(ctx *gin.Context) {
	var req dtos.CreateCompanyDTO
	if err := ctx.ShouldBind(&req); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "image is required", err.Error())
		return
	}

	uploadDir := "uploads/company/"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	filename := timestamp + "_" + file.Filename
	imagePath := filepath.Join(uploadDir, filename)

	if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to upload image", err.Error())
		return
	}

	data, err := c.service.Create(req, imagePath)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to create data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success create data", data)
}

func (c *CompanyController) Update(ctx *gin.Context) {
	var req dtos.UpdateCompanyDTO
	if err := ctx.ShouldBind(&req); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid data id", nil)
		return
	}

	var imagePath string
	file, err := ctx.FormFile("image")
	if err == nil {
		uploadDir := "uploads/company/"
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)

		filename := timestamp + "_" + file.Filename
		imagePath := filepath.Join(uploadDir, filename)

		if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
			helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to upload image", err.Error())
			return
		}
	}

	data, err := c.service.Update(req, imagePath, uint(id))
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to update data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success update data", data)
}

func (c *CompanyController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid data id", nil)
		return
	}

	err = c.service.Delete(uint(id))
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to delete data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success delete data", nil)
}
