package routes

import "github.com/gofiber/fiber"

var RegisterFeedRoutes = func(app fiber.App) {
	app.Get("/v1/api/users/:id", handlers.HandlerGetUserById)
	app.Get("/v1/api/posts/:id", handlers.HandlerGetPostById)
	app.Get("/v1/api/posts", handlers.HandlerGetPosts)
	app.Post("/v1/api/users", handlers.HandlerCreateUser)
	app.Post("/v1/api/login", handlers.HandlerUserLogin)
	app.Post("/v1/api/posts", handlers.HandlerCreatePost)
}
