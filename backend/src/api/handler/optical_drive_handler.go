package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Vigiatonet/PSU-Calculator/src/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/src/api/helper"
	"github.com/Vigiatonet/PSU-Calculator/src/config"
	"github.com/Vigiatonet/PSU-Calculator/src/services"
	"github.com/gin-gonic/gin"
)

type OpticalDriveHandler struct {
	service *services.OpticalDriveService
}

func NewOpticalDriveHandler(cfg *config.Config) *OpticalDriveHandler {
	s := services.NewOpticalDriveService(cfg)
	return &OpticalDriveHandler{
		service: s,
	}
}

// GetOpticalDriveById godoc
// @Summary Get OpticalDrive By Id
// @Description GetOpticalDriveById
// @Tags OpticalDrive
// @Accept json
// @produces json
// @Success 200 {object} helper.Response "GetOpticalDriveById response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/optical-drive/{id} [get]
// @Security AuthBearer
func (o *OpticalDriveHandler) GetOpticalDriveById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.ValidationError, false, errors.New("id must be a Uint")))
		return
	}
	res, err := o.service.GetOpticalDriveById(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, helper.Success, true))
}

// CreateOpticalDrive godoc
// @Summary Create a OpticalDrive
// @Description Create a OpticalDrive
// @Tags OpticalDrive
// @Accept json
// @produces json
// @Param Request body dto.CreateOpticalDriveRequest true "Create a OpticalDrive"
// @Success 201 {object} helper.Response "OpticalDrive response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/optical-drive/ [post]
// @Security AuthBearer
func (o *OpticalDriveHandler) CreateOpticalDrive(ctx *gin.Context) {
	req := &dto.CreateOpticalDriveRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.ValidationError, false, err))
		return
	}
	res, err := o.service.CreateOpticalDrive(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, helper.Success, true))
}

// UpdateOpticalDrive godoc
// @Summary Update OpticalDrive
// @Description Update OpticalDrive
// @Tags OpticalDrive
// @Accept json
// @produces json
// @Param Request body dto.UpdateOpticalDriveRequest true " Update OpticalDrive"
// @Success 200 {object} helper.Response "OpticalDrive response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/optical-drive/{id} [put]
// @Security AuthBearer
func (o *OpticalDriveHandler) UpdateOpticalDrive(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.ValidationError, false, errors.New("id must be a Uint")))
		return
	}
	req := &dto.UpdateOpticalDriveRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.ValidationError, false, err))
		return
	}
	res, err := o.service.UpdateOpticalDrive(ctx, req, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, helper.Success, true))

}

// DeleteOpticalDrive godoc
// @Summary Delete OpticalDrive
// @Description Delete OpticalDrive
// @Tags OpticalDrive
// @Accept json
// @produces json
// @Success 204 {object} helper.Response "OpticalDrive response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/optical-drive/{id} [delete]
// @Security AuthBearer
func (o *OpticalDriveHandler) DeleteOpticalDrive(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.ValidationError, false, errors.New("id must be a Uint")))
		return
	}
	err := o.service.DeleteOpticalDrive(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(gin.H{"Status": "Deleted"}, helper.Success, true))

}

// GetAllWithPagination godoc
// @Summary Get All With Pagination
// @Description Get All With Pagination
// @Tags OpticalDrive
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Get All With Pagination"
// @Success 201 {object} helper.Response "OpticalDrive response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/optical-drive/all/ [post]
// @Security AuthBearer
func (o *OpticalDriveHandler) GetAllWithPagination(ctx *gin.Context) {
	req := &dto.PaginationInputWithFilter{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.ValidationError, false, err))
		return
	}
	res, err := o.service.GetAllByFilter(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, helper.Success, true))
}
