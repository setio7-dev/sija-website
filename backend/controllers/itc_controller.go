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

type ItcController struct {
	service services.ItcService
}

func NewItcController() *ItcController {
	service := services.NewItcService()
	return &ItcController{service}
}

func (c *ItcController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to get data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success get all data", data)
}

func (c *ItcController) GetById(ctx *gin.Context) {
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

func (c *ItcController) Create(ctx *gin.Context) {
	var req dtos.CreateItcDTO
	if err := ctx.ShouldBind(&req); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "image is required", err.Error())
		return
	}

	uploadDir := "uploads/itc/"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	filename := timestamp + "_" + file.Filename
	filePath := filepath.Join(uploadDir, filename)

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to upload image", err.Error())
		return
	}

	data, err := c.service.Create(req, filePath)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to create data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success create data", data)
}

func (c *ItcController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid data id", nil)
		return
	}

	var req dtos.UpdateItcDTO
	if err := ctx.ShouldBind(&req); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	var imagePath string
	file, err := ctx.FormFile("image")
	if err == nil {
		uploadDir := "uploads/itc/"
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
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to update data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success update data", data)
}

func (c *ItcController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid data id", nil)
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to delete data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "data deleted successfully", nil)
}
