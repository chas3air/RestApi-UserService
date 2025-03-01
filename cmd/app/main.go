package main

import (
	"os"
	"os/signal"
	"syscall"
	"userservice/internal/app"
	"userservice/internal/storage/sqlite"
	"userservice/pkg/config"
	"userservice/pkg/logger"
)

func main() {
	config := config.MustLoad()

	logger := logger.SetupLogger(config.Env)

	_ = logger

	storage := sqlite.New(config.StoragePath)

	_ = storage

	application := app.New(logger, config.Address, storage)

	go func() {
		if err := application.Start(); err != nil {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
	logger.Info("Gracefully stopped")
}
