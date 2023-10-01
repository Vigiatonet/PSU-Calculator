package handler

import (
	"net/http"

	"github.com/Vigiatonet/PSU-Calculator/src/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/src/api/helper"
	"github.com/Vigiatonet/PSU-Calculator/src/config"
	"github.com/Vigiatonet/PSU-Calculator/src/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	s := services.NewUserService(cfg)
	return &UserHandler{
		service: s,
	}
}

func (uh *UserHandler) RegisterUser(ctx *gin.Context) {
	req := &dto.RegisterByUsername{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.ValidationError, false, err))
		return
	}
	err = uh.service.RegisterByUsername(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.NotFoundError, false, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse("created", helper.Success, true))
}

func (uh *UserHandler) LoginUser(ctx *gin.Context) {
	req := &dto.LoginByUsername{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.ValidationError, false, err))
		return
	}
	result, err := uh.service.LoginByUsername(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.NotFoundError, false, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(result, helper.Success, true))
}

func (us *UserHandler) RefreshToken(ctx *gin.Context) {
	req := &dto.RefreshTokenDTO{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.ValidationError, false, err))
		return
	}
	tk, err := us.service.Token.ValidateRefreshToken(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.InternalError, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(tk, helper.Success, true))
}

// TODO: add logout after auth middleWare
