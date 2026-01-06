package logger

import (
	"log/slog"
	"os"
)

type Options struct {
	Env string
}

func New(opts Options) *slog.Logger {
	var h slog.Handler
	if opts.Env == "prod" {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}

	l :=slog.New(h)
	slog.SetDefault(l)
	return l 
}
