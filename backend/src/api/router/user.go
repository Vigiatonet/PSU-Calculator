package router

import (
	"github.com/Vigiatonet/PSU-Calculator/src/api/handler"
	"github.com/Vigiatonet/PSU-Calculator/src/api/middleware"
	"github.com/Vigiatonet/PSU-Calculator/src/config"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUserHandler(cfg)
	r.POST("/register", h.RegisterUser)
	r.POST("/login", h.LoginUser)
	r.POST("/refresh", middleware.Authentication(cfg), h.RefreshToken)
}
