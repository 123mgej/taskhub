package router
import (
	"github.com/gin-gonic/gin"
	"taskhub/internal/handle"
)

func New() *gin.Engine {
	r :=gin.New()

	// middlewares
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	r.GET("/healthz",handle.Healthz)
	return r
}