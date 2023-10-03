package handler

import (
	"net/http"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/api/helper"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/services"
	"github.com/gin-gonic/gin"
)

type GraphicHandler struct {
	service *services.GraphicService
}

func NewGraphicHandler(cfg *config.Config) *GraphicHandler {
	s := services.NewGraphicService(cfg)
	return &GraphicHandler{
		service: s,
	}
}

// GetGraphicById godoc
// @Summary Get Graphic By Id
// @Description GetGraphicById
// @Tags Graphic
// @Accept json
// @produces json
// @Success 200 {object} helper.Response "GetGraphicById response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/graphic/get/{id} [get]
// @Security AuthBearer
func (o *GraphicHandler) GetGraphicById(ctx *gin.Context) {

	GetById[dto.GraphicResponse](ctx, o.service.GetGraphicById)

}

// CreateGraphic godoc
// @Summary Create a Graphic
// @Description Create a Graphic
// @Tags Graphic
// @Accept json
// @produces json
// @Param Request body dto.CreateGraphicRequest true "Create a Graphic"
// @Success 201 {object} helper.Response "Graphic response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/graphic/create [post]
// @Security AuthBearer
func (o *GraphicHandler) CreateGraphic(ctx *gin.Context) {
	Create[dto.CreateGraphicRequest, dto.GraphicResponse](ctx, o.service.CreateGraphic)
}

// UpdateGraphic godoc
// @Summary Update Graphic
// @Description Update Graphic
// @Tags Graphic
// @Accept json
// @produces json
// @Param Request body dto.UpdateGraphicRequest true " Update Graphic"
// @Success 200 {object} helper.Response "Graphic response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/graphic/update/{id} [put]
// @Security AuthBearer
func (o *GraphicHandler) UpdateGraphic(ctx *gin.Context) {
	Update[dto.UpdateGraphicRequest, dto.GraphicResponse](ctx, o.service.UpdateGraphic)
}

// DeleteGraphic godoc
// @Summary Delete Graphic
// @Description Delete Graphic
// @Tags Graphic
// @Accept json
// @produces json
// @Success 204 {object} helper.Response "Graphic response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/graphic/delete/{id} [delete]
// @Security AuthBearer
func (o *GraphicHandler) DeleteGraphic(ctx *gin.Context) {

	Delete(ctx, o.service.DeleteGraphic)
}

// GetAllWithPagination godoc
// @Summary Get All With Pagination
// @Description Get All With Pagination
// @Tags Graphic
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Get All With Pagination"
// @Success 201 {object} helper.Response "Graphic response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/graphic/all/ [post]
// @Security AuthBearer
func (o *GraphicHandler) GetAllWithPagination(ctx *gin.Context) {
	GetByFilter[dto.GraphicResponse](ctx, o.service.GetAllByFilter)
}

// GetAllWithBrand godoc
// @Summary Get All With Brand filter
// @Description Get All With Brand filter
// @Tags Graphic
// @Accept json
// @produces json
// @Success 201 {object} helper.Response "Graphic response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/graphic/brand/{brand} [get]
// @Security AuthBearer
func (o *GraphicHandler) GetAllWithBrand(ctx *gin.Context) {
	brand := ctx.Params.ByName("brand")

	res, err := o.service.GetAllByBrand(brand)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, helper.Success, true))
}
