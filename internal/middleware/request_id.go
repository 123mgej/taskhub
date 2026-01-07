package middleware

import (
	requestid "taskhub/internal/pkg/request_id"

	"github.com/gin-gonic/gin"
)

const (
	HeaderRequestID = "X-Request-ID"
	CtxRequestIDKey = "request_id"
)

func RequestID() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		// 获取 request id 全局变量
		rid :=ctx.GetHeader(HeaderRequestID)
		// request id 全局变量不存在，再重新生成 request id
		if rid == "" {
			rid =  requestid.New()
		}

		ctx.Set(CtxRequestIDKey,rid)
		ctx.Header(HeaderRequestID,rid)

		ctx.Next()
	}
}

func GetRequestID (c *gin.Context) string {
	if c == nil {
		return ""
	}
	if v,ok :=c.Get(CtxRequestIDKey); ok {
		if s,ok :=v.(string); ok {
			return  s
		}
	}
	return ""
}