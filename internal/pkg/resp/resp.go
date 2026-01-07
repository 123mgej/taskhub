package resp

import (
	"net/http"
	"taskhub/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	RequestID string    `json:"request_id,omitempty`
}

func Ok(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "ok",
		Data:    data,
		RequestID: middleware.GetRequestID(c),
	})
}

func Failed(c *gin.Context, code int, messgae string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: messgae,
		RequestID: middleware.GetRequestID(c),
	})
}
