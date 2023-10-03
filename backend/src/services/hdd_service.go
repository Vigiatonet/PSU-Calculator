package services

import (
	"context"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/data/models"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
)

type HardDriveService struct {
	base *BaseService[models.HardDrive, dto.UpdateHardDriveRequest, dto.CreateHardDriveRequest, dto.HardDriveResponse]
}

func NewHardDriveService(cfg *config.Config) *HardDriveService {
	db := db.GetDB()
	logger := logging.NewLogger(cfg)
	base := &BaseService[models.HardDrive, dto.UpdateHardDriveRequest, dto.CreateHardDriveRequest, dto.HardDriveResponse]{
		DB:     db,
		Logger: logger,
	}
	return &HardDriveService{
		base: base,
	}
}

func (o *HardDriveService) CreateHardDrive(ctx context.Context, req *dto.CreateHardDriveRequest) (*dto.HardDriveResponse, error) {
	return o.base.Create(ctx, req)
}

func (o *HardDriveService) GetHardDriveById(ctx context.Context, id int) (*dto.HardDriveResponse, error) {
	return o.base.GetById(ctx, id)
}

func (o *HardDriveService) UpdateHardDrive(ctx context.Context, req *dto.UpdateHardDriveRequest, id int) (*dto.HardDriveResponse, error) {
	return o.base.Update(ctx, req, id)
}

func (o *HardDriveService) DeleteHardDrive(ctx context.Context, id int) error {
	return o.base.Delete(ctx, id)
}

func (o *HardDriveService) GetAllByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.HardDriveResponse], error) {
	return o.base.GetByFilter(ctx, req)
}
