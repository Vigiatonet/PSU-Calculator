package router

import (
	"github.com/Vigiatonet/PSU-Calculator/api/handler"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/gin-gonic/gin"
)

func OpticalDriveRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewOpticalDriveHandler(cfg)
	r.GET("/:id", h.GetOpticalDriveById)
	r.PUT("/:id", h.UpdateOpticalDrive)
	r.DELETE("/:id", h.DeleteOpticalDrive)
	r.POST("/", h.CreateOpticalDrive)
	r.POST("/all/", h.GetAllWithPagination)

}
