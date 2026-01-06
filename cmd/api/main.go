package main

import (
	"log/slog"
	"taskhub/internal/config"
	"taskhub/internal/pkg/logger"
	"taskhub/internal/router"
)

func main() {
	r := router.New()

	config := config.Load()

	addr := ":" + config.Port

	log := logger.New(logger.Options{Env: config.Env})
	log.Info("server starting", slog.String("env", config.Env), slog.String("addr", addr))
	slog.String("taskhub api listening on %s\n", addr)

	if err := r.Run(addr); err != nil {
		log.Error("server failed ", slog.Any("err",err))
	}
}
