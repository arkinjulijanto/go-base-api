package boot

import (
	"github.com/arkinjulijanto/go-base-api/apps/api/routes"
	"github.com/arkinjulijanto/go-base-api/boot/base"
	"github.com/arkinjulijanto/go-base-api/config"
	"github.com/arkinjulijanto/go-base-api/internal/handlers"
	"github.com/arkinjulijanto/go-base-api/internal/utils/gin_util"
	"github.com/arkinjulijanto/go-base-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func InitApi() error {
	config.LoadEnv()
	conf := config.GetEnv()
	r := gin.Default()
	r.Use(gin_util.RequestIDMiddleware)

	logger.Init(conf)
	err := base.InitDB(conf)
	if err != nil {
		return err
	}

	h := handlers.InitHandlers()

	routes.InitRoutes(r, h)

	err = r.Run(conf.APP_PORT)
	if err != nil {
		return err
	}

	return nil
}
