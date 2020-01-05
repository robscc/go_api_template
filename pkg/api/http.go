package api

import (
	"go_api_template/pkg/api/controller"
	_ "go_api_template/pkg/api/docs"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Go API Template
// @version 1.0.0
// @description Go API Template

// @contact.name Chenchuan.Song
// @contact.email robscc269@gmail.com
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY

// ServeAPI ...
func ServeAPI(cfg *Config) error {

	r := gin.New()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, If-Modified-Since",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	r.GET("/ping", controller.Ping)
	r.GET("/test_auth", controller.TestAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r.Run(cfg.BindAddr)
}
