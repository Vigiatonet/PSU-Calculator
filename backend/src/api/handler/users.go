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

// RegisterUser godoc
// @Summary RegisterUser
// @Description register users with username and password
// @Tags Users
// @Accept json
// @produces json
// @Param Request body dto.RegisterByUsername true "RegisterByUsername"
// @Success 201 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/register [post]
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

// LoginUser godoc
// @Summary LoginUser
// @Description Login users with username and password
// @Tags Users
// @Accept json
// @produces json
// @Param Request body dto.LoginByUsername true "LoginUser"
// @Success 201 {object} helper.Response "response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/login [post]
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

// RefreshToken godoc
// @Summary RefreshToken
// @Description Get a new AccessToken With Refresh token (key rotation is on)
// @Tags Users
// @Accept json
// @produces json
// @Param Request body dto.RefreshTokenDTO true "Get a new AccessToken With Refresh token"
// @Success 201 {object} helper.Response{result=dto.TokenDetail} "AccessToken response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/refresh [post]
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
