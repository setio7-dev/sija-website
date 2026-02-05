package controllers

import (
	"sijaku-hebat/dtos"
	"sijaku-hebat/helpers"
	"sijaku-hebat/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ModuleController struct {
	service services.ModuleService
}

func NewModuleController() *ModuleController {
	service := services.NewModuleService()
	return &ModuleController{service}
}

func (c *ModuleController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "failed to get data",
			"error":   err.Error(),
		})
		return
	}

	helpers.SuccessResponse(ctx, "success get all data", data)
}

func (c *ModuleController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(ctx, 400, "invalid data id", nil)
		return
	}

	data, err := c.service.GetById(uint(id))
	if err != nil {
		helpers.ErrorResponse(ctx, 404, "data not found", err.Error())
		return
	}
	helpers.SuccessResponse(ctx, "success get data by id", data)
}

func (c *ModuleController) Create(ctx *gin.Context) {
	var req dtos.CreateModuleDTO
	if err := ctx.ShouldBind(&req); err != nil {
		helpers.ErrorResponse(ctx, 400, "invalid request body", err.Error())
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		helpers.ErrorResponse(ctx, 400, "image is required", err.Error())
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		helpers.ErrorResponse(ctx, 400, "file is required", err.Error())
		return
	}

	uploadImageDir := "uploads/module/images"
	uploadFileDir := "uploads/module/files"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	imageFilename := timestamp + "_" + image.Filename
	imagePath := uploadImageDir + "/" + imageFilename
	if err := ctx.SaveUploadedFile(image, imagePath); err != nil {
		helpers.ErrorResponse(ctx, 500, "failed to upload image", err.Error())
		return
	}

	fileFilename := timestamp + "_" + file.Filename
	filePath := uploadFileDir + "/" + fileFilename
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		helpers.ErrorResponse(ctx, 500, "failed to upload file", err.Error())
		return
	}

	data, err := c.service.Create(req, imagePath, filePath)
	if err != nil {
		helpers.ErrorResponse(ctx, 422, "failed to create data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "data created successfully", data)
}

func (c *ModuleController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(ctx, 400, "invalid data id", nil)
		return
	}

	var req dtos.UpdateModuleDTO
	if err := ctx.ShouldBind(&req); err != nil {
		helpers.ErrorResponse(ctx, 400, "invalid request body", err.Error())
		return
	}

	var imagePath string
	image, err := ctx.FormFile("image")
	if err == nil {
		uploadImageDir := "uploads/module/images"
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)

		imageFilename := timestamp + "_" + image.Filename
		imagePath = uploadImageDir + "/" + imageFilename

		if err := ctx.SaveUploadedFile(image, imagePath); err != nil {
			helpers.ErrorResponse(ctx, 500, "failed to upload image", err.Error())
			return
		}

	}

	var filePath string
	file, err := ctx.FormFile("file")
	if err == nil {
		uploadFileDir := "uploads/module/files"
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)

		fileFilename := timestamp + "_" + file.Filename
		filePath = uploadFileDir + "/" + fileFilename

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			helpers.ErrorResponse(ctx, 500, "failed to upload file", err.Error())
			return
		}
	}

	data, err := c.service.Update(req, imagePath, filePath, uint(id))
	if err != nil {
		helpers.ErrorResponse(ctx, 500, "failed to update data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "data updated successfully", data)
}

func (c *ModuleController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(ctx, 400, "invalid data id", nil)
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		helpers.ErrorResponse(ctx, 500, "failed to delete data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "data deleted successfully", nil)
}
