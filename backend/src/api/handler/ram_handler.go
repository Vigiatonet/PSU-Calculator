package handler

import (
	"net/http"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/api/helper"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/services"
	"github.com/gin-gonic/gin"
)

type RamModelHandler struct {
	service *services.RamModelService
}

func NewRamModelHandler(cfg *config.Config) *RamModelHandler {
	s := services.NewRamModelService(cfg)
	return &RamModelHandler{
		service: s,
	}
}

// GetRamModelById godoc
// @Summary Get Ram Model By Id
// @Description Get Ram Model By Id
// @Tags RamModel
// @Accept json
// @produces json
// @Success 200 {object} helper.Response "GetRamModelById response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ram/get/{id} [get]
// @Security AuthBearer
func (o *RamModelHandler) GetRamModelById(ctx *gin.Context) {

	GetById[dto.RamModelResponse](ctx, o.service.GetRamModelById)

}

// CreateRamModel godoc
// @Summary Create a Ram Model
// @Description Create a Ram Model
// @Tags RamModel
// @Accept json
// @produces json
// @Param Request body dto.CreateRamModelRequest true "Create a RamModel"
// @Success 201 {object} helper.Response "RamModel response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ram/create [post]
// @Security AuthBearer
func (o *RamModelHandler) CreateRamModel(ctx *gin.Context) {
	Create[dto.CreateRamModelRequest, dto.RamModelResponse](ctx, o.service.CreateRamModel)
}

// UpdateRamModel godoc
// @Summary Update RamModel
// @Description Update RamModel
// @Tags RamModel
// @Accept json
// @produces json
// @Param Request body dto.UpdateRamModelRequest true " Update RamModel"
// @Success 200 {object} helper.Response "RamModel response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ram/update/{id} [put]
// @Security AuthBearer
func (o *RamModelHandler) UpdateRamModel(ctx *gin.Context) {
	Update[dto.UpdateRamModelRequest, dto.RamModelResponse](ctx, o.service.UpdateRamModel)
}

// DeleteRamModel godoc
// @Summary Delete RamModel
// @Description Delete RamModel from data base
// @Tags RamModel
// @Accept json
// @produces json
// @Success 204 {object} helper.Response "RamModel response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ram/delete/{id} [delete]
// @Security AuthBearer
func (o *RamModelHandler) DeleteRamModel(ctx *gin.Context) {

	Delete(ctx, o.service.DeleteRamModel)
}

// GetAllWithPagination godoc
// @Summary Get All With Pagination
// @Description Get All With Pagination
// @Tags RamModel
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Get All With Pagination"
// @Success 201 {object} helper.Response "RamModel response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ram/all/ [post]
// @Security AuthBearer
func (o *RamModelHandler) GetAllWithPagination(ctx *gin.Context) {
	GetByFilter[dto.RamModelResponse](ctx, o.service.GetAllByFilter)
	/*
		{
		"filter": {
			"type": {
			"form": "ddr5",
			"type": "contains"
			}
		},
		"page_number": 0,
		"page_size": 0,
		"sort": [
			{
			"colId": "ramSize",
			"string": "desc"
			}
		]
		}
	*/
}

// GetAllByType godoc
// @Summary Get All based on ddr4 or ddr5 or ddr3
// @Description  Get All based on ddr4 or ddr5 or ddr3
// @Tags RamModel
// @Accept json
// @produces json
// @Success 201 {object} helper.Response "RamModel response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ram/type/{type} [get]
// @Security AuthBearer
func (o *RamModelHandler) GetAllByType(ctx *gin.Context) {
	ramType := ctx.Params.ByName("type")
	res, err := o.service.GetByRamType(ramType)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, helper.Success, true))

}
