package api

import (
	"fmt"

	"github.com/Vigiatonet/PSU-Calculator/src/config"
	"github.com/Vigiatonet/PSU-Calculator/src/docs"
	"github.com/Vigiatonet/PSU-Calculator/src/pkg/logging"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	var log = logging.NewLogger(cfg)

	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger()) // TODO: add custom logger

	swaggerInit(cfg, r)
	registerRoutes(r)

	err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatal(err, logging.Internal, logging.Api, "run failed ", nil)
		return
	}
}

func registerRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		_ = v1
	}
}

func swaggerInit(cfg *config.Config, r *gin.Engine) {

	docs.SwaggerInfo.Title = "psu calculator"
	docs.SwaggerInfo.Description = "api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", cfg.Server.Port)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
