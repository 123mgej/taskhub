package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		latency := time.Since(start)
		status := ctx.Writer.Status()

		path := ctx.FullPath()
		rid := GetRequestID(ctx)

		attr := []slog.Attr{
			slog.Int64("eslaped time", latency.Milliseconds()),
			slog.Int("status", status),
			slog.String("path", path),
			slog.String("request id", rid),
			slog.String("method", ctx.Request.Method),
		}

		if len(ctx.Errors) > 0 && ctx.Errors.Last() != nil {
			attr = append(attr, slog.String("eror", ctx.Errors.Last().Error()))
		}

		log := slog.Default()
		switch {
		case status >= 500:
			log.LogAttrs(ctx, slog.LevelError, "access", attr...)

		case status >= 400:
			log.LogAttrs(ctx, slog.LevelWarn, "access", attr...)

		default:
			log.LogAttrs(ctx, slog.LevelInfo, "access", attr...)

		}
	}
}
