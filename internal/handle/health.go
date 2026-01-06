package handle

import (
	"taskhub/internal/pkg/resp"

	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context) {
	resp.Ok(c, gin.H{"status": "ok"})
}
