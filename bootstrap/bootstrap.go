package bootstrap

import (
	"context"

	"behealth-api/api/controllers"
	"behealth-api/api/middlewares"
	"behealth-api/api/repository"
	"behealth-api/api/routes"
	"behealth-api/api/services"
	"behealth-api/infrastructure"
	"behealth-api/seeds"

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
	seeds.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.RequestHandler,
	routes routes.Routes,
	env infrastructure.Env,
	logger infrastructure.Logger,
	database infrastructure.Database,
	migrations infrastructure.Migrations,
	seeds seeds.Seeds,

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

				logger.Zap.Info("🖇️  Seting up route ....")
				routes.Setup()

				logger.Zap.Info(" 🌱 Seeding data ......")
				seeds.Run()

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
