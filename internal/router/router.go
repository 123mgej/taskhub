package router

import (
	"taskhub/internal/config"
	"taskhub/internal/handle"
	"taskhub/internal/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func New(cfg config.Config) *gin.Engine {
	r :=gin.New()

	// middlewares
	r.Use(middleware.RequestID())
	r.Use(middleware.AccessLog())
	r.Use(gin.Recovery())


	if cfg.Env != "prod" {
			pprof.Register(r,"/debug/pprof")

	}

	r.GET("/healthz",handle.Healthz)
	return r
}