package services

import (
	"context"
	"errors"
	"strings"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/data/models"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
)

type RamModelService struct {
	base *BaseService[models.Ram, dto.UpdateRamModelRequest, dto.CreateRamModelRequest, dto.RamModelResponse]
}

func NewRamModelService(cfg *config.Config) *RamModelService {
	db := db.GetDB()
	logger := logging.NewLogger(cfg)
	base := &BaseService[models.Ram, dto.UpdateRamModelRequest, dto.CreateRamModelRequest, dto.RamModelResponse]{
		DB:     db,
		Logger: logger,
	}
	return &RamModelService{
		base: base,
	}
}

func (o *RamModelService) CreateRamModel(ctx context.Context, req *dto.CreateRamModelRequest) (*dto.RamModelResponse, error) {
	if !strings.HasPrefix(req.Type, "ddr") {
		return nil, errors.New("invalid type for ram should start with ddr ")
	}
	return o.base.Create(ctx, req)
}

func (o *RamModelService) GetRamModelById(ctx context.Context, id int) (*dto.RamModelResponse, error) {
	return o.base.GetById(ctx, id)
}

func (o *RamModelService) UpdateRamModel(ctx context.Context, req *dto.UpdateRamModelRequest, id int) (*dto.RamModelResponse, error) {
	return o.base.Update(ctx, req, id)
}

func (o *RamModelService) DeleteRamModel(ctx context.Context, id int) error {
	return o.base.Delete(ctx, id)
}

func (o *RamModelService) GetAllByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.RamModelResponse], error) {
	return o.base.GetByFilter(ctx, req)
}

func (o *RamModelService) GetByRamType(ramType string) (*[]dto.RamModelResponse, error) {
	if !strings.HasPrefix(ramType, "ddr") {
		return nil, errors.New("invalid type for ram should start with ddr ")
	}
	var result []dto.RamModelResponse
	err := o.base.DB.Model(&models.Ram{}).Where("type = ?", ramType).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
