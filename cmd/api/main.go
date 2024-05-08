package main

import (
	"github.com/arkinjulijanto/go-base-api/boot"
	"github.com/arkinjulijanto/go-base-api/pkg/logger"
)

func main() {
	err := boot.InitApi()
	if err != nil {
		logger.LogFatal("failed to initialize API", err, nil)
	}
}
