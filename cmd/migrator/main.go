package main

import (
	"github.com/arkinjulijanto/go-base-api/boot/base"
	"github.com/arkinjulijanto/go-base-api/config"
	"github.com/arkinjulijanto/go-base-api/internal/models"
	"github.com/arkinjulijanto/go-base-api/pkg/logger"
)

func main() {
	config.LoadEnv()
	conf := config.GetEnv()

	logger.Init(conf)
	err := base.InitDB(conf)
	if err != nil {
		logger.LogFatal("failed to init db", err, nil)
	}

	db := base.GetDBConn()

	db.AutoMigrate(
		&models.User{},
	)

	logger.LogInfo("success migration", nil)
}
