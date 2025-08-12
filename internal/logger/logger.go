package logger

import (
	"github.com/M-kos/wb_level0/internal/config"
	"log/slog"
	"os"
)

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
	Warn(msg string, args ...any)
	Debug(msg string, args ...any)
}

type logger struct {
	Log *slog.Logger
}

func NewLogger(config *config.Config) Logger {
	mapLevels := map[string]slog.Level{
		"DEBUG": slog.LevelDebug,
		"INFO":  slog.LevelInfo,
		"WARN":  slog.LevelWarn,
		"ERROR": slog.LevelError,
	}

	level := slog.LevelDebug

	if config != nil && config.LogLevel != "" {
		level = mapLevels[config.LogLevel]
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	return &logger{
		Log: log,
	}
}

func (l *logger) Info(msg string, args ...any) {
	l.Log.Info(msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.Log.Error(msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.Log.Warn(msg, args...)
}

func (l *logger) Debug(msg string, args ...any) {
	l.Log.Debug(msg, args...)
}
