package routes

import (
	"am_chat/controllers"
	"am_chat/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {

	user := r.Group("/api")

	user.Get("/", middlewares.Auth, controllers.Index)
	user.Get("/:idrec", controllers.Show)
	user.Post("/", middlewares.Auth, controllers.Create)
	user.Put("/:idrec", controllers.Update)
	user.Delete("/:idrec", controllers.Delete)
}
