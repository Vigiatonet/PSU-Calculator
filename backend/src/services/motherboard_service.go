package services

import (
	"context"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/data/models"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
)

type MotherboardService struct {
	base *BaseService[models.Motherboard, dto.UpdateMotherboardRequest, dto.CreateMotherboardRequest, dto.MotherboardResponse]
}

func NewMotherboardService(cfg *config.Config) *MotherboardService {
	db := db.GetDB()
	logger := logging.NewLogger(cfg)
	base := &BaseService[models.Motherboard, dto.UpdateMotherboardRequest, dto.CreateMotherboardRequest, dto.MotherboardResponse]{
		DB:     db,
		Logger: logger,
	}
	return &MotherboardService{
		base: base,
	}
}

func (o *MotherboardService) CreateMotherboard(ctx context.Context, req *dto.CreateMotherboardRequest) (*dto.MotherboardResponse, error) {
	return o.base.Create(ctx, req)
}

func (o *MotherboardService) GetMotherboardById(ctx context.Context, id int) (*dto.MotherboardResponse, error) {
	return o.base.GetById(ctx, id)
}

func (o *MotherboardService) UpdateMotherboard(ctx context.Context, req *dto.UpdateMotherboardRequest, id int) (*dto.MotherboardResponse, error) {
	return o.base.Update(ctx, req, id)
}

func (o *MotherboardService) DeleteMotherboard(ctx context.Context, id int) error {
	return o.base.Delete(ctx, id)
}

func (o *MotherboardService) GetAllByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.MotherboardResponse], error) {
	return o.base.GetByFilter(ctx, req)
}
