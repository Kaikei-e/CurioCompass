package logger

import (
	"context"
	"log/slog"
	"os"
)

func SetUpLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"structured logging is enabled",
	)

	slog.Info("Logger is successfully set up! FeedFetcher is ready to go!")
}
