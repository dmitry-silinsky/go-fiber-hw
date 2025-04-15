package logger

import (
	"dmitry-silinsky/go-fiber-hw/config"
	"fmt"
	"log/slog"
	"os"
)

func NewLogger(logConfig *config.LogConfig) (*slog.Logger, error) {
	var handler slog.Handler

	handlerOptions := &slog.HandlerOptions{
		Level: slog.Level(logConfig.Level),
	}

	switch logConfig.Format {
	case config.LogFormatJSON:
		handler = slog.NewJSONHandler(os.Stdout, handlerOptions)
	case config.LogFormatText:
		handler = slog.NewTextHandler(os.Stdout, handlerOptions)
	default:
		return nil, fmt.Errorf("Формат логов \"%s\" не поддерживается", logConfig.Format)
	}

	return slog.New(handler), nil
}
