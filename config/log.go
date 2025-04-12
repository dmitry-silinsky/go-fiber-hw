package config

import (
	"strings"
)

type LogFormat string

const (
	LogFormatJSON LogFormat = "json"
	LogFormatText LogFormat = "text"
)

type LogConfig struct {
	Format LogFormat
	Level int
}

func NewLogConfig() *LogConfig {
	formatVal := strings.ToLower(
		getString("LOG_FORMAT", "text"),
	)

	return &LogConfig{
		Format: LogFormat(formatVal),
		Level: getInt("LOG_LEVEL", 0),
	}
}
