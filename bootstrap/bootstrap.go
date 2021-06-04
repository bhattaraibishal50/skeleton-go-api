package bootstrap

import (
	"context"

	"behealth-api/api/controllers"
	"behealth-api/api/middlewares"
	"behealth-api/api/repository"
	"behealth-api/api/routes"
	"behealth-api/api/services"
	"behealth-api/infrastructure"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	middlewares.Module,
	routes.Module,
	infrastructure.Module,
	services.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.RequestHandler,
	routes routes.Routes,
	env infrastructure.Env,
	logger infrastructure.Logger,
	database infrastructure.Database,
	middlewares middlewares.Middlewares,
	migrations infrastructure.Migrations,

) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("---- Starting Application ----")
			logger.Zap.Info("------------------------------")
			logger.Zap.Info("-------- Be-health API -------")
			logger.Zap.Info("------------------------------")

			logger.Zap.Info(" 🚌 Migrating DB Schema .......")
			migrations.Migrate()

			conn.SetMaxOpenConns(10)
			go func() {
				routes.Setup()
				if env.ServerPort == "" {
					handler.Gin.Run()
				} else {
					handler.Gin.Run(":" + env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info(" 🛑 Stopping Application .....")
			conn.Close()
			return nil
		},
	})
}
