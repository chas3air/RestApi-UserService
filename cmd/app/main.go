package main

import (
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

	// router

	// run
}
