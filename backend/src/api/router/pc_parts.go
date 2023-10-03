package router

import (
	"github.com/Vigiatonet/PSU-Calculator/api/handler"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/gin-gonic/gin"
)

func HardDriveRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewHardDriveHandler(cfg)
	r.GET("/get/:id", h.GetHardDriveById)
	r.PUT("/update/:id", h.UpdateHardDrive)
	r.DELETE("/delete/:id", h.DeleteHardDrive)
	r.POST("/create", h.CreateHardDrive)
	r.POST("/all/", h.GetAllWithPagination)
}

func RamModelRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewRamModelHandler(cfg)
	r.GET("/get/:id", h.GetRamModelById)
	r.GET("/type/:type", h.GetAllByType)
	r.PUT("/update/:id", h.UpdateRamModel)
	r.DELETE("/delete/:id", h.DeleteRamModel)
	r.POST("/create", h.CreateRamModel)
	r.POST("/all/", h.GetAllWithPagination)
}

func MotherboardRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewMotherboardHandler(cfg)
	r.GET("/get/:id", h.GetMotherboardById)
	r.PUT("/update/:id", h.UpdateMotherboard)
	r.DELETE("/delete/:id", h.DeleteMotherboard)
	r.POST("/create", h.CreateMotherboard)
	r.POST("/all/", h.GetAllWithPagination)
}

func GraphicRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewGraphicHandler(cfg)
	r.GET("/get/:id", h.GetGraphicById)
	r.PUT("/update/:id", h.UpdateGraphic)
	r.DELETE("/delete/:id", h.DeleteGraphic)
	r.POST("/create", h.CreateGraphic)
	r.GET("/brand/:brand/", h.GetAllWithBrand)
	r.POST("/all/", h.GetAllWithPagination)
}

func CpuRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCpuHandler(cfg)
	r.GET("/get/:id", h.GetCpuById)
	r.PUT("/update/:id", h.UpdateCpu)
	r.DELETE("/delete/:id", h.DeleteCpu)
	r.POST("/create", h.CreateCpu)
	r.GET("/brand/:brand/", h.GetAllByBrand)
	r.POST("/all/", h.GetAllWithPagination)
}

func SsdRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewSsdHandler(cfg)
	r.GET("/get/:id", h.GetSsdById)
	r.PUT("/update/:id", h.UpdateSsd)
	r.DELETE("/delete/:id", h.DeleteSsd)
	r.POST("/create", h.CreateSsd)
	r.POST("/all/", h.GetAllWithPagination)
}

func PowerRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPowerHandler(cfg)
	r.POST("/calc/", h.CalculatePower)
}
