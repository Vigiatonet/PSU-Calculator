package handler

import (
	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/services"
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
// @Router /v1/optical-drive/get/{id} [get]
// @Security AuthBearer
func (o *OpticalDriveHandler) GetOpticalDriveById(ctx *gin.Context) {

	GetById[dto.OpticalDriveResponse](ctx, o.service.GetOpticalDriveById)

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
// @Router /v1/optical-drive/create [post]
// @Security AuthBearer
func (o *OpticalDriveHandler) CreateOpticalDrive(ctx *gin.Context) {
	Create[dto.CreateOpticalDriveRequest, dto.OpticalDriveResponse](ctx, o.service.CreateOpticalDrive)
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
// @Router /v1/optical-drive/update/{id} [put]
// @Security AuthBearer
func (o *OpticalDriveHandler) UpdateOpticalDrive(ctx *gin.Context) {
	Update[dto.UpdateOpticalDriveRequest, dto.OpticalDriveResponse](ctx, o.service.UpdateOpticalDrive)
}

// DeleteOpticalDrive godoc
// @Summary Delete OpticalDrive
// @Description Delete OpticalDrive
// @Tags OpticalDrive
// @Accept json
// @produces json
// @Success 204 {object} helper.Response "OpticalDrive response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/optical-drive/delete/{id} [delete]
// @Security AuthBearer
func (o *OpticalDriveHandler) DeleteOpticalDrive(ctx *gin.Context) {

	Delete(ctx, o.service.DeleteOpticalDrive)
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
	GetByFilter[dto.OpticalDriveResponse](ctx, o.service.GetAllByFilter)
}
