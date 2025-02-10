package logger

import (
	"log"
	"log/slog"
	"os"
)

func SetUpLogger(env string) *slog.Logger {
	switch env {
	case "prod":
		return slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelInfo}))
	case "local":
		return slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{Level: slog.LevelDebug}))
	default:
		log.Fatalf("invalid env %s", env)
	}
	return nil
}
