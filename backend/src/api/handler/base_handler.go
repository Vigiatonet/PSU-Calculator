package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/api/helper"
	"github.com/gin-gonic/gin"
)

// FIXME: if these generic functions break there is no way i can debug this. (But is saves time)
// There is nothing we can do -Napoleon

func GetById[Tr any](ctx *gin.Context, fn func(ctx context.Context, id int) (*Tr, error)) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.ValidationError, false, errors.New("id must be a Uint")))
		return
	}

	res, err := fn(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, helper.Success, true))
}

func Create[Tc, Tr any](ctx *gin.Context, fn func(ctx context.Context, req *Tc) (*Tr, error)) {
	req := new(Tc)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.ValidationError, false, err))
		return
	}
	res, err := fn(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, helper.Success, true))
}

func Update[Tu, Tr any](ctx *gin.Context, fn func(ctx context.Context, req *Tu, id int) (*Tr, error)) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.ValidationError, false, errors.New("id must be a Uint")))
		return
	}
	req := new(Tu)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.ValidationError, false, err))
		return
	}
	res, err := fn(ctx, req, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, helper.Success, true))
}

func Delete(ctx *gin.Context, fn func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if id <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.ValidationError, false, errors.New("id must be a Uint")))
		return
	}
	err := fn(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusNoContent, helper.GenerateBaseResponse(gin.H{"Status": "Deleted"}, helper.Success, true))

}

func GetByFilter[Tr any](ctx *gin.Context, fn func(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[Tr], error)) {
	req := &dto.PaginationInputWithFilter{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(helper.ValidationError, false, err))
		return
	}
	res, err := fn(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, helper.Error, false, err))
		return
	}
	ctx.JSON(http.StatusOK, helper.GenerateBaseResponse(res, helper.Success, true))
}
