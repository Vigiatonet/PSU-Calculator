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

type GraphicService struct {
	base *BaseService[models.Graphic, dto.UpdateGraphicRequest, dto.CreateGraphicRequest, dto.GraphicResponse]
}

func NewGraphicService(cfg *config.Config) *GraphicService {
	db := db.GetDB()
	logger := logging.NewLogger(cfg)
	base := &BaseService[models.Graphic, dto.UpdateGraphicRequest, dto.CreateGraphicRequest, dto.GraphicResponse]{
		DB:     db,
		Logger: logger,
		Preload: []Preloads{
			{
				name: "GpuBrand",
			},
		},
	}
	return &GraphicService{
		base: base,
	}
}

func (o *GraphicService) CreateGraphic(ctx context.Context, req *dto.CreateGraphicRequest) (*dto.GraphicResponse, error) {
	return o.base.Create(ctx, req)
}

func (o *GraphicService) GetGraphicById(ctx context.Context, id int) (*dto.GraphicResponse, error) {
	return o.base.GetById(ctx, id)
}

func (o *GraphicService) UpdateGraphic(ctx context.Context, req *dto.UpdateGraphicRequest, id int) (*dto.GraphicResponse, error) {
	return o.base.Update(ctx, req, id)
}

func (o *GraphicService) DeleteGraphic(ctx context.Context, id int) error {
	return o.base.Delete(ctx, id)
}

func (o *GraphicService) GetAllByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.GraphicResponse], error) {
	return o.base.GetByFilter(ctx, req)
}

func (o *GraphicService) GetAllByBrand(brand string) (*[]dto.GraphicResponse, error) {
	var result *[]dto.GraphicResponse
	brandName := strings.Title(strings.ToLower(brand))
	var brandId int

	err := o.base.DB.Model(&models.GpuBrand{}).Select("id").Where("name = ?", brandName).Find(&brandId).Error
	if err != nil {
		return nil, err
	}

	err = o.base.DB.Model(&models.Graphic{}).Where("gpu_brand_id = ?", brandId).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
