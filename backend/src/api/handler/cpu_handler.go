package handler

import (
	"net/http"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/api/helper"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/services"
	"github.com/gin-gonic/gin"
)

type CpuHandler struct {
	service *services.CpuService
}

func NewCpuHandler(cfg *config.Config) *CpuHandler {
	s := services.NewCpuService(cfg)
	return &CpuHandler{
		service: s,
	}
}

// GetCpuById godoc
// @Summary Get Cpu By Id
// @Description GetCpuById
// @Tags Cpu
// @Accept json
// @produces json
// @Success 200 {object} helper.Response "GetCpuById response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/cpu/get/{id} [get]
// @Security AuthBearer
func (o *CpuHandler) GetCpuById(ctx *gin.Context) {

	GetById[dto.CpuResponse](ctx, o.service.GetCpuById)

}

// CreateCpu godoc
// @Summary Create a Cpu
// @Description Create a Cpu
// @Tags Cpu
// @Accept json
// @produces json
// @Param Request body dto.CreateCpuRequest true "Create a Cpu"
// @Success 201 {object} helper.Response "Cpu response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/cpu/create [post]
// @Security AuthBearer
func (o *CpuHandler) CreateCpu(ctx *gin.Context) {
	Create[dto.CreateCpuRequest, dto.CpuResponse](ctx, o.service.CreateCpu)
}

// UpdateCpu godoc
// @Summary Update Cpu
// @Description Update Cpu
// @Tags Cpu
// @Accept json
// @produces json
// @Param Request body dto.UpdateCpuRequest true " Update Cpu"
// @Success 200 {object} helper.Response "Cpu response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/cpu/update/{id} [put]
// @Security AuthBearer
func (o *CpuHandler) UpdateCpu(ctx *gin.Context) {
	Update[dto.UpdateCpuRequest, dto.CpuResponse](ctx, o.service.UpdateCpu)
}

// DeleteCpu godoc
// @Summary Delete Cpu
// @Description Delete Cpu
// @Tags Cpu
// @Accept json
// @produces json
// @Success 204 {object} helper.Response "Cpu response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/cpu/delete/{id} [delete]
// @Security AuthBearer
func (o *CpuHandler) DeleteCpu(ctx *gin.Context) {
	Delete(ctx, o.service.DeleteCpu)
}

// GetAllWithPagination godoc
// @Summary Get All With Pagination
// @Description Get All With Pagination
// @Tags Cpu
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Get All With Pagination"
// @Success 201 {object} helper.Response "Cpu response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/cpu/all/ [post]
// @Security AuthBearer
func (o *CpuHandler) GetAllWithPagination(ctx *gin.Context) {
	GetByFilter[dto.CpuResponse](ctx, o.service.GetAllByFilter)
}

// GetAllByBrand godoc
// @Summary Get All By Brand
// @Description Get All By Brand
// @Tags Cpu
// @Accept json
// @produces json
// @Success 201 {object} helper.Response "Cpu response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/cpu/brand/{brand} [get]
// @Security AuthBearer
func (o *CpuHandler) GetAllByBrand(ctx *gin.Context) {
	brand := ctx.Params.ByName("brand")
	res, err := o.service.GetAllByBrand(brand)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, helper.Success, true))
}
