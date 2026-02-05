package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/priyanshu334/go-bend/internal/config"
	"github.com/priyanshu334/go-bend/internal/db"
	"github.com/priyanshu334/go-bend/internal/logger"
	"github.com/priyanshu334/go-bend/internal/routes"
	"go.uber.org/zap"
)

func main() {
	config.Load()
	logger.Init(config.Cfg.AppEnv)
	if err := db.Connect(); err != nil {
		log.Fatal("db connection failed", err)
	}
	if err := db.Migrate(); err != nil {
		log.Fatal("migration failed", err)
	}

	app := fiber.New()

	routes.Register(app)
	app.Get("/health", func(c *fiber.Ctx) error {
		logger.Log.Info("health check hit")
		return c.JSON(fiber.Map{"status": "ok"})
	})

	logger.Log.Info("🚀 server started",
		zap.String("port", config.Cfg.AppPort),
	)

	app.Listen(":" + config.Cfg.AppPort)

}
