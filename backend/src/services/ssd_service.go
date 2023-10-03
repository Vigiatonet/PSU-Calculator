package services

import (
	"context"

	"github.com/Vigiatonet/PSU-Calculator/src/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/src/config"
	"github.com/Vigiatonet/PSU-Calculator/src/data/db"
	"github.com/Vigiatonet/PSU-Calculator/src/data/models"
	"github.com/Vigiatonet/PSU-Calculator/src/pkg/logging"
)

type SsdService struct {
	base *BaseService[models.Ssd, dto.UpdateSsdRequest, dto.CreateSsdRequest, dto.SsdResponse]
}

func NewSsdService(cfg *config.Config) *SsdService {
	db := db.GetDB()
	logger := logging.NewLogger(cfg)
	base := &BaseService[models.Ssd, dto.UpdateSsdRequest, dto.CreateSsdRequest, dto.SsdResponse]{
		DB:     db,
		Logger: logger,
	}
	return &SsdService{
		base: base,
	}
}

func (o *SsdService) CreateSsd(ctx context.Context, req *dto.CreateSsdRequest) (*dto.SsdResponse, error) {
	return o.base.Create(ctx, req)
}

func (o *SsdService) GetSsdById(ctx context.Context, id int) (*dto.SsdResponse, error) {
	return o.base.GetById(ctx, id)
}

func (o *SsdService) UpdateSsd(ctx context.Context, req *dto.UpdateSsdRequest, id int) (*dto.SsdResponse, error) {
	return o.base.Update(ctx, req, id)
}

func (o *SsdService) DeleteSsd(ctx context.Context, id int) error {
	return o.base.Delete(ctx, id)
}

func (o *SsdService) GetAllByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.SsdResponse], error) {
	return o.base.GetByFilter(ctx, req)
}
