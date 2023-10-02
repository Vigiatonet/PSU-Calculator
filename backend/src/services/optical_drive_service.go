package services

import (
	"context"

	"github.com/Vigiatonet/PSU-Calculator/src/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/src/config"
	"github.com/Vigiatonet/PSU-Calculator/src/data/db"
	"github.com/Vigiatonet/PSU-Calculator/src/data/models"
	"github.com/Vigiatonet/PSU-Calculator/src/pkg/logging"
)

type OpticalDriveService struct {
	base *BaseService[models.OpticalDrive, dto.UpdateOpticalDriveRequest, dto.CreateOpticalDriveRequest, dto.OpticalDriveResponse]
}

func NewOpticalDriveService(cfg *config.Config) *OpticalDriveService {
	db := db.GetDB()
	logger := logging.NewLogger(cfg)
	base := &BaseService[models.OpticalDrive, dto.UpdateOpticalDriveRequest, dto.CreateOpticalDriveRequest, dto.OpticalDriveResponse]{
		DB:     db,
		Logger: logger,
	}
	return &OpticalDriveService{
		base: base,
	}
}

func (o *OpticalDriveService) CreateOpticalDrive(ctx context.Context, req *dto.CreateOpticalDriveRequest) (*dto.OpticalDriveResponse, error) {
	return o.base.Create(ctx, req)
}

func (o *OpticalDriveService) GetOpticalDriveById(ctx context.Context, id int) (*dto.OpticalDriveResponse, error) {
	return o.base.GetById(ctx, id)
}

func (o *OpticalDriveService) UpdateOpticalDrive(ctx context.Context, req *dto.UpdateOpticalDriveRequest, id int) (*dto.OpticalDriveResponse, error) {
	return o.base.Update(ctx, req, id)
}

func (o *OpticalDriveService) DeleteOpticalDrive(ctx context.Context, id int) error {
	return o.base.Delete(ctx, id)
}

func (o *OpticalDriveService) GetAllByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.OpticalDriveResponse], error) {
	return o.base.GetByFilter(ctx, req)
}
