package router

import (
	"taskhub/internal/handle"
	"taskhub/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r :=gin.New()

	// middlewares
	r.Use(middleware.RequestID())
	r.Use(middleware.AccessLog())
	r.Use(gin.Recovery())


	r.GET("/healthz",handle.Healthz)
	return r
}