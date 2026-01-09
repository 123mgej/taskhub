package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"taskhub/internal/pkg/token"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "user_id"

func Auth(tm *token.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := ctx.GetHeader("Authorization")
		fmt.Printf("Authorization=%q\n",h)
		// 从 http header 中提取 Authorization的内容
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"code":40101,"message":"missing token",
			})
			return
		}

		raw := strings.TrimPrefix(h, "Bearer ")
		// 从 token 中拿到 uid
		uid, err := tm.Parse(raw)
		if err != nil || uid == 0 {
			
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"code":40101,"message":"invalid token",
			})
			return
		}

		ctx.Set(CtxUserIDKey, uid)
		ctx.Next()

	}
}

func GetUserID(c *gin.Context) uint64 {
	if v, ok := c.Get(CtxUserIDKey); ok {
		if id, ok := v.(uint64); ok {
			return id
		}

	}
	return 0
}
