package handler

import (
	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/services"
	"github.com/gin-gonic/gin"
)

type HardDriveHandler struct {
	service *services.HardDriveService
}

func NewHardDriveHandler(cfg *config.Config) *HardDriveHandler {
	s := services.NewHardDriveService(cfg)
	return &HardDriveHandler{
		service: s,
	}
}

// GetHardDriveById godoc
// @Summary Get HardDrive By Id
// @Description GetHardDriveById
// @Tags HardDrive
// @Accept json
// @produces json
// @Success 200 {object} helper.Response "GetHardDriveById response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/hdd/get/{id} [get]
// @Security AuthBearer
func (o *HardDriveHandler) GetHardDriveById(ctx *gin.Context) {

	GetById[dto.HardDriveResponse](ctx, o.service.GetHardDriveById)

}

// CreateHardDrive godoc
// @Summary Create a HardDrive
// @Description Create a HardDrive
// @Tags HardDrive
// @Accept json
// @produces json
// @Param Request body dto.CreateHardDriveRequest true "Create a HardDrive"
// @Success 201 {object} helper.Response "HardDrive response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/hdd/create [post]
// @Security AuthBearer
func (o *HardDriveHandler) CreateHardDrive(ctx *gin.Context) {
	Create[dto.CreateHardDriveRequest, dto.HardDriveResponse](ctx, o.service.CreateHardDrive)
}

// UpdateHardDrive godoc
// @Summary Update HardDrive
// @Description Update HardDrive
// @Tags HardDrive
// @Accept json
// @produces json
// @Param Request body dto.UpdateHardDriveRequest true " Update HardDrive"
// @Success 200 {object} helper.Response "HardDrive response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/hdd/update/{id} [put]
// @Security AuthBearer
func (o *HardDriveHandler) UpdateHardDrive(ctx *gin.Context) {
	Update[dto.UpdateHardDriveRequest, dto.HardDriveResponse](ctx, o.service.UpdateHardDrive)
}

// DeleteHardDrive godoc
// @Summary Delete HardDrive
// @Description Delete HardDrive
// @Tags HardDrive
// @Accept json
// @produces json
// @Success 204 {object} helper.Response "HardDrive response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/hdd/delete/{id} [delete]
// @Security AuthBearer
func (o *HardDriveHandler) DeleteHardDrive(ctx *gin.Context) {

	Delete(ctx, o.service.DeleteHardDrive)
}

// GetAllWithPagination godoc
// @Summary Get All With Pagination
// @Description Get All With Pagination
// @Tags HardDrive
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Get All With Pagination"
// @Success 201 {object} helper.Response "HardDrive response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/hdd/all/ [post]
// @Security AuthBearer
func (o *HardDriveHandler) GetAllWithPagination(ctx *gin.Context) {
	GetByFilter[dto.HardDriveResponse](ctx, o.service.GetAllByFilter)
}
