package boot

import (
	"github.com/arkinjulijanto/go-base-api/apps/api/routes"
	"github.com/arkinjulijanto/go-base-api/boot/base"
	"github.com/arkinjulijanto/go-base-api/config"
	"github.com/arkinjulijanto/go-base-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitApi() error {
	conf := config.GetEnv()
	r := gin.Default()

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
