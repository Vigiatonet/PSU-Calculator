package api

import (
	"fmt"

	"github.com/Vigiatonet/PSU-Calculator/api/middleware"
	"github.com/Vigiatonet/PSU-Calculator/api/router"
	"github.com/Vigiatonet/PSU-Calculator/api/validators"
	"github.com/Vigiatonet/PSU-Calculator/config"
	"github.com/Vigiatonet/PSU-Calculator/docs"
	"github.com/Vigiatonet/PSU-Calculator/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	var log = logging.NewLogger(cfg)
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors(cfg), middleware.CustomLogger(log), middleware.Limiter())

	swaggerInit(cfg, r)
	registerRoutes(r, cfg)
	registerValidators()

	err := r.Run(fmt.Sprintf(":%d", cfg.Server.InternalPort))
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

		opticalDrive := v1.Group("/optical-drive", middleware.Authentication(cfg))
		router.OpticalDriveRouter(opticalDrive, cfg)

		hdd := v1.Group("/hdd", middleware.Authentication(cfg))
		router.HardDriveRouter(hdd, cfg)

		ram := v1.Group("/ram", middleware.Authentication(cfg))
		router.RamModelRouter(ram, cfg)

		mb := v1.Group("/motherboard", middleware.Authentication(cfg))
		router.MotherboardRouter(mb, cfg)

		gpu := v1.Group("/gpu", middleware.Authentication(cfg))
		router.GraphicRouter(gpu, cfg)

		cpu := v1.Group("/cpu", middleware.Authentication(cfg))
		router.CpuRouter(cpu, cfg)

		ssd := v1.Group("/ssd", middleware.Authentication(cfg))
		router.SsdRouter(ssd, cfg)

		pw := v1.Group("/power", middleware.Authentication(cfg))
		router.PowerRouter(pw, cfg)
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
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", cfg.Server.ExternalPort)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
