package providers

import (
	"log/slog"
	"os"
	"strings"

	"github.com/dpnam2112/go-backend-template/internal/config"
	"go.uber.org/fx"
)

func ProvideLogger(cfg *config.Config) *slog.Logger {
	var level slog.Level
	switch strings.ToLower(cfg.LogLevel) {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return slog.New(handler)
}

var LoggerModule = fx.Module("logger",
	fx.Provide(ProvideLogger),
)

