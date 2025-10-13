package logger

import (
	"github.com/devvdark0/book-library/internal/config"
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func InitLogger(cfg *config.Config) Logger {
	var handler slog.Handler
	switch cfg.Env {
	case "prod":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	case "dev":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	return Logger{
		logger: slog.New(handler),
	}
}

func (l Logger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l Logger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l Logger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l Logger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
