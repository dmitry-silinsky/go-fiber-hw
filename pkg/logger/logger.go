package logger

import (
	"dmitry-silinsky/go-fiber-hw/config"
	"log/slog"
	"os"
)

func NewLogger(logConfig *config.LogConfig) *slog.Logger {
	var handler slog.Handler

	handlerOptions := &slog.HandlerOptions{
		Level: slog.Level(logConfig.Level),
	}

	switch logConfig.Format {
	case config.LogFormatJSON:
		handler = slog.NewJSONHandler(os.Stdout, handlerOptions)
	case config.LogFormatText:
		handler = slog.NewTextHandler(os.Stdout, handlerOptions)
	}

	return slog.New(handler)
}
