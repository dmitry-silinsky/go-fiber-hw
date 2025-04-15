package main

import (
	"dmitry-silinsky/go-fiber-hw/config"
	"dmitry-silinsky/go-fiber-hw/internal/pages"
	"dmitry-silinsky/go-fiber-hw/pkg/logger"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.Load()

	appConfig := config.NewAppConfig()
	logConfig := config.NewLogConfig()

	logger, err := logger.NewLogger(logConfig)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	// Подключаем логирование через slog
	app.Use(slogfiber.New(logger))
	// Подключаем recovery middleware для восстановления
	// работы приложения при возникновании ошибки
	app.Use(recover.New())

	pages.NewHandler(app)

	app.Listen(appConfig.ListenAddr)
}
