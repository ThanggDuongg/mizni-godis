package main

import (
	"context"
	"log/slog"
	"mizni-godis/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := server.New(":6379", logger)
	if err := srv.Start(ctx); err != nil {
		logger.Error("server exited", "err", err)
		os.Exit(1)
	}

	logger.Info("goodbye")
}
