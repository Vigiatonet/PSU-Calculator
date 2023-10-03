package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/api/helper"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/constants"
	"github.com/Vigiatonet/PSU-Calculator/services"
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
// @Success 200 {object} helper.Response "response"
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
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(result, helper.Success, true))
}

// RefreshToken godoc
// @Summary RefreshToken
// @Description Get a new AccessToken With Refresh token (key rotation is on)
// @Tags Users
// @Accept json
// @produces json
// @Param Request body dto.RefreshTokenDTO true "Get a new AccessToken With Refresh token"
// @Success 200 {object} helper.Response{result=dto.TokenDetail} "AccessToken response"
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

// LogoutUser godoc
// @Summary LogoutUser
// @Description Get a new logout user and expire JwtToken
// @Tags Users
// @Accept json
// @produces json
// @Success 418 {object} helper.Response "AccessToken response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/logout [post]
func (uh *UserHandler) LogoutUser(ctx *gin.Context) {
	tokenHeader := ctx.GetHeader(constants.AuthenTicationHeaderKey)
	token := strings.Split(tokenHeader, " ")[1]
	err := services.AddToBlacklist(token, uh.service.Token.Cfg.JWT.AccessTokenExpireDuration*time.Minute)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError("field to logout", helper.InternalError, false, err))
		return
	}
	ctx.JSON(http.StatusTeapot, helper.GenerateBaseResponse("loggedOut successfully", helper.Success, true))

}

// ShowUserDetail godoc
// @Summary ShowUserDetail
// @Description show user details
// @Tags Users
// @Accept json
// @produces json
// @Success 200 {object} helper.Response "UserDetail response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/users/profile [post]
func (uh *UserHandler) ShowUserDetail(ctx *gin.Context) {
	res, err := uh.service.ShowUser(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(nil, helper.ValidationError, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, helper.Success, true))
}
