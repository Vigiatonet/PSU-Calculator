package handler

import (
	"net/http"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/api/helper"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/services"
	"github.com/gin-gonic/gin"
)

type PowerHandler struct {
	Service *services.PowerCalculateService
}

func NewPowerHandler(cfg *config.Config) *PowerHandler {
	s := services.NewPowerCalculateService(cfg)
	return &PowerHandler{
		Service: s,
	}
}

// CalculatePower godoc
// @Summary Calculate Power
// @Description Calculate Power
// @Tags power
// @Accept json
// @produces json
// @Param Request body dto.CalculatePowerRequest true " Calculate Power"
// @Success 200 {object} helper.Response "Ssd response"
// @Failure 400 {object} helper.Response "Bad request"
// @Router /v1/ssd/power [post]
// @Security AuthBearer
func (ph *PowerHandler) CalculatePower(ctx *gin.Context) {
	req := dto.CalculatePowerRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.Error, false, err))
		return
	}
	res := ph.Service.CalculatePower(&req)
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, helper.Success, true))
}
