package controllers

import (
	"net/http"
	"sijaku-hebat/dtos"
	"sijaku-hebat/helpers"
	"sijaku-hebat/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController() *UserController {
	service := services.NewUserService()
	return &UserController{service}
}

func (c *UserController) GetAll(ctx *gin.Context) {
	data, err := c.service.GetAll()
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to get data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success get all data", data)
}

func (c *UserController) GetById(ctx *gin.Context) {
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

func (c *UserController) Login(ctx *gin.Context) {
	var req dtos.LoginUserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	token, err := c.service.Login(req)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusUnauthorized, "login failed", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "login successful", gin.H{"token": token})
}

func (c *UserController) Create(ctx *gin.Context) {
	var req dtos.CreateUserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	data, err := c.service.Create(req)
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusInternalServerError, "failed to create data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success create data", data)
}

func (c *UserController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid data id", nil)
		return
	}

	var req dtos.UpdateUserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helpers.ErrorResponse(ctx, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	data, err := c.service.Update(req, uint(id))
	if err != nil {
		helpers.ErrorResponse(ctx, http.StatusUnprocessableEntity, "failed to update data", err.Error())
		return
	}

	helpers.SuccessResponse(ctx, "success update data", data)
}

func (c *UserController) Delete(ctx *gin.Context) {
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
