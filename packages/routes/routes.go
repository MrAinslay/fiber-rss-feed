package routes

import (
	"github.com/MrAinslay/fiber-rss-feed/packages/handlers"
	"github.com/MrAinslay/fiber-rss-feed/packages/middleware"
	"github.com/gofiber/fiber"
)

var RegisterFeedRoutes = func(app fiber.App) {
	app.Get("/v1/api/users/:api_key", handlers.HandlerGetUserByKey)
	app.Get("/v1/api/posts/:id", handlers.HandlerGetPostById)
	app.Get("/v1/api/feeds/:id", handlers.HnadlerGetFeedById)
	app.Get("/v1/api/posts", handlers.HandlerGetPosts)
	app.Get("/v1/api/feeds", handlers.HandlerGetFeeds)
	app.Post("/v1/api/feeds", middleware.MiddlewareAuth(handlers.HandlerCreateFeed))
	app.Post("/v1/api/users", handlers.HandlerCreateUser)
	app.Post("/v1/api/login", handlers.HandlerUserLogin)
	app.Post("/v1/api/posts", handlers.HandlerCreatePost)
}
