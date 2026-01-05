package handle
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"message":"ok",
	})
}