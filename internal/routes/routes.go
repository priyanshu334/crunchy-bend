package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/priyanshu334/go-bend/internal/db"
	"github.com/priyanshu334/go-bend/internal/moduels/user"
)

func Register(app *fiber.App) {
	api := app.Group("/api")

	UserRepo := user.NewRepository(db.DB)
	userService := user.NewService(UserRepo)
	userHandler := user.NewHandler(userService)
	user.RegisterRoutes(api, userHandler)
}
