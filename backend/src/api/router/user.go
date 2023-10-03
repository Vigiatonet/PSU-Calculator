package router

import (
	"github.com/Vigiatonet/PSU-Calculator/api/handler"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)
	r.POST("/register", h.RegisterUser)
	r.POST("/login", h.LoginUser)
	r.POST("/refresh", h.RefreshToken)
}
