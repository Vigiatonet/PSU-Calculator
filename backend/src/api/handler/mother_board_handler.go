package handler

import (
	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/services"
	"github.com/gin-gonic/gin"
)

type MotherboardHandler struct {
	service *services.MotherboardService
}

func NewMotherboardHandler(cfg *config.Config) *MotherboardHandler {
	s := services.NewMotherboardService(cfg)
	return &MotherboardHandler{
		service: s,
	}
}

// GetMotherboardById godoc
// @Summary Get Motherboard By Id
// @Description GetMotherboardById
// @Tags Motherboard
// @Accept json
// @produces json
// @Success 200 {object} helper.Response "GetMotherboardById response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/motherboard/get/{id} [get]
// @Security AuthBearer
func (o *MotherboardHandler) GetMotherboardById(ctx *gin.Context) {

	GetById[dto.MotherboardResponse](ctx, o.service.GetMotherboardById)

}

// CreateMotherboard godoc
// @Summary Create a Motherboard
// @Description Create a Motherboard
// @Tags Motherboard
// @Accept json
// @produces json
// @Param Request body dto.CreateMotherboardRequest true "Create a Motherboard"
// @Success 201 {object} helper.Response "Motherboard response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/motherboard/create [post]
// @Security AuthBearer
func (o *MotherboardHandler) CreateMotherboard(ctx *gin.Context) {
	Create[dto.CreateMotherboardRequest, dto.MotherboardResponse](ctx, o.service.CreateMotherboard)
}

// UpdateMotherboard godoc
// @Summary Update Motherboard
// @Description Update Motherboard
// @Tags Motherboard
// @Accept json
// @produces json
// @Param Request body dto.UpdateMotherboardRequest true "Update Motherboard"
// @Success 200 {object} helper.Response "Motherboard response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/motherboard/update/{id} [put]
// @Security AuthBearer
func (o *MotherboardHandler) UpdateMotherboard(ctx *gin.Context) {
	Update[dto.UpdateMotherboardRequest, dto.MotherboardResponse](ctx, o.service.UpdateMotherboard)
}

// DeleteMotherboard godoc
// @Summary Delete Motherboard
// @Description Delete Motherboard
// @Tags Motherboard
// @Accept json
// @produces json
// @Success 204 {object} helper.Response "Motherboard response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/motherboard/delete/{id} [delete]
// @Security AuthBearer
func (o *MotherboardHandler) DeleteMotherboard(ctx *gin.Context) {

	Delete(ctx, o.service.DeleteMotherboard)
}

// GetAllWithPagination godoc
// @Summary Get All With Pagination
// @Description Get All With Pagination
// @Tags Motherboard
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Get All With Pagination"
// @Success 201 {object} helper.Response "Motherboard response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/motherboard/all/ [post]
// @Security AuthBearer
func (o *MotherboardHandler) GetAllWithPagination(ctx *gin.Context) {
	GetByFilter[dto.MotherboardResponse](ctx, o.service.GetAllByFilter)
}
