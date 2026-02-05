package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/priyanshu334/go-bend/internal/db"
	"github.com/priyanshu334/go-bend/internal/moduels/anime"
	"github.com/priyanshu334/go-bend/internal/moduels/auth"
	"github.com/priyanshu334/go-bend/internal/moduels/image"
	"github.com/priyanshu334/go-bend/internal/moduels/user"
)

func Register(app *fiber.App) {
	api := app.Group("/api")

	UserRepo := user.NewRepository(db.DB)
	userService := user.NewService(UserRepo)
	userHandler := user.NewHandler(userService)
	authRepo := auth.NewRepository(UserRepo)

	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	imageRepo := image.NewRepository(db.DB)
	imageService := image.NewService(imageRepo)
	imageHandler := image.NewHandler(imageService)

	animeRepo := anime.NewRepository(db.DB)
	animeService := anime.NewService(animeRepo)
	animeHandler := anime.NewHandler(animeService)

	user.RegisterRoutes(api, userHandler)
	auth.RegisterRoutes(api, authHandler)
	image.RegisterRoutes(api, imageHandler)
	anime.RegisterRoutes(api, animeHandler)

}
