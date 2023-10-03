package services

import (
	"context"
	"strings"

	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/data/models"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
)

type CpuService struct {
	base *BaseService[models.Cpu, dto.UpdateCpuRequest, dto.CreateCpuRequest, dto.CpuResponse]
}

func NewCpuService(cfg *config.Config) *CpuService {
	db := db.GetDB()
	logger := logging.NewLogger(cfg)
	base := &BaseService[models.Cpu, dto.UpdateCpuRequest, dto.CreateCpuRequest, dto.CpuResponse]{
		DB:     db,
		Logger: logger,
	}
	return &CpuService{
		base: base,
	}
}

func (o *CpuService) CreateCpu(ctx context.Context, req *dto.CreateCpuRequest) (*dto.CpuResponse, error) {
	return o.base.Create(ctx, req)
}

func (o *CpuService) GetCpuById(ctx context.Context, id int) (*dto.CpuResponse, error) {
	return o.base.GetById(ctx, id)
}

func (o *CpuService) UpdateCpu(ctx context.Context, req *dto.UpdateCpuRequest, id int) (*dto.CpuResponse, error) {
	return o.base.Update(ctx, req, id)
}

func (o *CpuService) DeleteCpu(ctx context.Context, id int) error {
	return o.base.Delete(ctx, id)
}

func (o *CpuService) GetAllByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.CpuResponse], error) {
	return o.base.GetByFilter(ctx, req)
}

func (o *CpuService) GetAllByBrand(brand string) (*[]dto.CpuResponse, error) {
	brandName := strings.Title(strings.ToLower(brand))
	var brandId int
	err := o.base.DB.Model(&models.CpuBrand{}).Select("id").Where("name = ?", brandName).First(&brandId).Error
	if err != nil {
		return nil, err
	}
	var result *[]dto.CpuResponse
	err = o.base.DB.Model(&models.Cpu{}).Where("cpu_brand_id = ?", brandId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
