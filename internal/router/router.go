package router

import (
	"taskhub/internal/config"
	"taskhub/internal/handle"
	"taskhub/internal/middleware"
	"taskhub/internal/pkg/resp"
	"taskhub/internal/pkg/token"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(cfg config.Config,db *gorm.DB) *gin.Engine {
	r :=gin.New()

	// middlewares
	r.Use(middleware.RequestID())
	r.Use(middleware.AccessLog())
	r.Use(gin.Recovery())
	


	if cfg.Env != "prod" {
			pprof.Register(r,"/debug/pprof")

	}

	tm:=token.New(cfg.JWTSecret,cfg.JWTExpireMinutes)
	ah:=&handle.AuthHandler{DB: db,TM: tm}
	api:=r.Group("/api/v1")
	api.POST("/auth/register",ah.Register)
	api.POST("/auth/login",ah.Login)
	apiAuth:=api.Group("")
	apiAuth.Use(middleware.Auth(tm))
	apiAuth.GET("/me",func(ctx *gin.Context) {
		uid:=middleware.GetUserID(ctx)
		resp.Ok(ctx,gin.H{"user_id":uid})
	})



	r.GET("/healthz",handle.Healthz)
	return r
}