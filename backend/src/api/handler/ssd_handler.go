package handler

import (
	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/services"
	"github.com/gin-gonic/gin"
)

type SsdHandler struct {
	service *services.SsdService
}

func NewSsdHandler(cfg *config.Config) *SsdHandler {
	s := services.NewSsdService(cfg)
	return &SsdHandler{
		service: s,
	}
}

// GetSsdById godoc
// @Summary Get Ssd By Id
// @Description GetSsdById
// @Tags Ssd
// @Accept json
// @produces json
// @Success 200 {object} helper.Response "GetSsdById response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ssd/get/{id} [get]
// @Security AuthBearer
func (o *SsdHandler) GetSsdById(ctx *gin.Context) {

	GetById[dto.SsdResponse](ctx, o.service.GetSsdById)

}

// CreateSsd godoc
// @Summary Create a Ssd
// @Description Create a Ssd
// @Tags Ssd
// @Accept json
// @produces json
// @Param Request body dto.CreateSsdRequest true "Create a Ssd"
// @Success 201 {object} helper.Response "Ssd response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ssd/create [post]
// @Security AuthBearer
func (o *SsdHandler) CreateSsd(ctx *gin.Context) {
	Create[dto.CreateSsdRequest, dto.SsdResponse](ctx, o.service.CreateSsd)
}

// UpdateSsd godoc
// @Summary Update Ssd
// @Description Update Ssd
// @Tags Ssd
// @Accept json
// @produces json
// @Param Request body dto.UpdateSsdRequest true " Update Ssd"
// @Success 200 {object} helper.Response "Ssd response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ssd/update/{id} [put]
// @Security AuthBearer
func (o *SsdHandler) UpdateSsd(ctx *gin.Context) {
	Update[dto.UpdateSsdRequest, dto.SsdResponse](ctx, o.service.UpdateSsd)
}

// DeleteSsd godoc
// @Summary Delete Ssd
// @Description Delete Ssd
// @Tags Ssd
// @Accept json
// @produces json
// @Success 204 {object} helper.Response "Ssd response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ssd/delete/{id} [delete]
// @Security AuthBearer
func (o *SsdHandler) DeleteSsd(ctx *gin.Context) {
	Delete(ctx, o.service.DeleteSsd)
}

// GetAllWithPagination godoc
// @Summary Get All With Pagination
// @Description Get All With Pagination
// @Tags Ssd
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Get All With Pagination"
// @Success 200 {object} helper.Response "Ssd response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ssd/all/ [post]
// @Security AuthBearer
func (o *SsdHandler) GetAllWithPagination(ctx *gin.Context) {
	GetByFilter[dto.SsdResponse](ctx, o.service.GetAllByFilter)
}
