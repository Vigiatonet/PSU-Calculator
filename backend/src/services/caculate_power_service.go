package services

import (
	"github.com/Vigiatonet/PSU-Calculator/api/dto"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/data/db"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
	"gorm.io/gorm"
)

type PowerCalculateService struct {
	DB       *gorm.DB
	Logger   logging.Logger
	Cfg      *config.Config
	Preloads []Preloads
}

func NewPowerCalculateService(cfg *config.Config) *PowerCalculateService {
	l := logging.NewLogger(cfg)
	db := db.GetDB()
	return &PowerCalculateService{
		DB:     db,
		Logger: l,
		Cfg:    cfg,
		Preloads: []Preloads{
			{name: ""},
		},
	}
}

func (p *PowerCalculateService) CalculatePower(req *dto.CalculatePowerRequest) *dto.CalculatePowerResponse {
	var powerNeeded float64

	defaultCounts := map[*int]int{
		&req.RamCount:          1,
		&req.GraphicCount:      1,
		&req.SsdCount:          1,
		&req.HardDriveCount:    1,
		&req.OpticalDriveCount: 1,
	}

	for countPtr, defaultCount := range defaultCounts {
		if *countPtr == 0 {
			*countPtr = defaultCount
		}
	}

	powerNeeded = req.CpuPower +
		float64(req.RamCount)*req.RamPower +
		float64(req.GraphicCount)*req.GraphicPower +
		float64(req.SsdCount)*req.SsdPower +
		float64(req.HardDriveCount)*req.HardDrivePower +
		float64(req.OpticalDriveCount)*req.OpticalDrivePower +
		req.MotherboardPower

	return &dto.CalculatePowerResponse{TotalPowerConsumption: powerNeeded}

}

// func (p *PowerCalculateService) CalculatePowerById()
