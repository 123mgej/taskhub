package main

import (
	"log/slog"
	"taskhub/internal/config"
	"taskhub/internal/pkg/db"
	"taskhub/internal/pkg/logger"
	"taskhub/internal/router"
	"taskhub/internal/model"
)

func main() {

	
	config := config.Load()
	r := router.New(config)


	addr := ":" + config.Port

	log := logger.New(logger.Options{Env: config.Env})
	log.Info("server starting", slog.String("env", config.Env), slog.String("addr", addr))
	slog.String("taskhub api listening on %s\n", addr)

gdb, err := db.Open(db.Options{DSN: config.DB_DSN, Env: config.Env})
if err != nil {
  log.Error("db connect failed", slog.Any("err", err))
  return
}
if err := gdb.AutoMigrate(&model.User{}, &model.Task{}); err != nil {
  log.Error("db migrate failed", slog.Any("err", err))
  return
}
log.Info("db connected")


	if err := r.Run(addr); err != nil {
		log.Error("server failed ", slog.Any("err",err))
	}
}
