package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matiolsz/go-and-react/controllers"
	"github.com/matiolsz/go-and-react/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
}
