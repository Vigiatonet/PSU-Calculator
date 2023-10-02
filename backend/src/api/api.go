package api

import (
	"fmt"

	"github.com/Vigiatonet/PSU-Calculator/src/api/middleware"
	"github.com/Vigiatonet/PSU-Calculator/src/api/router"
	"github.com/Vigiatonet/PSU-Calculator/src/api/validators"
	"github.com/Vigiatonet/PSU-Calculator/src/config"
	"github.com/Vigiatonet/PSU-Calculator/src/docs"
	"github.com/Vigiatonet/PSU-Calculator/src/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	var log = logging.NewLogger(cfg)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors(cfg), middleware.CustomLogger(log), middleware.Limiter())

	swaggerInit(cfg, r)
	registerRoutes(r, cfg)
	registerValidators()

	err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatal(err, logging.Internal, logging.Api, "run failed ", nil)
		return
	}
}

func registerRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		users := v1.Group("/users")
		router.UserRouter(users, cfg)
	}
}

func registerValidators() {
	log := logging.NewLogger(config.GetConfig())
	vld, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := vld.RegisterValidation("password", validators.ValidatePassword, true)
		if err != nil {
			log.Error(err, logging.General, logging.Api, "cant add validator password", nil)
			return
		}
		err = vld.RegisterValidation("email", validators.ValidateEmail, true)
		if err != nil {
			log.Error(err, logging.General, logging.Api, "cant add validator email", nil)
			return
		}
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
