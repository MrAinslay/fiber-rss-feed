package routes

import (
	"github.com/MrAinslay/fiber-rss-feed/packages/handlers"
	"github.com/MrAinslay/fiber-rss-feed/packages/middleware"
	"github.com/gofiber/fiber"
)

var RegisterFeedRoutes = func(app *fiber.App) {
	app.Get("/v1/api/users/:api_key", handlers.HandlerGetUserByKey)
	app.Get("/v1/api/posts/:id", handlers.HandlerGetPostById)
	app.Get("/v1/api/feeds/:id", handlers.HnadlerGetFeedById)
	app.Get("/v1/api/posts", handlers.HandlerGetPosts)
	app.Get("/v1/api/feeds", handlers.HandlerGetFeeds)
	app.Get("/v1/api/feed-follows", middleware.MiddlewareAuth(handlers.HandlerGetUserFeedFollows))
	app.Post("/v1/api/feeds", middleware.MiddlewareAuth(handlers.HandlerCreateFeed))
	app.Post("/v1/api/users", handlers.HandlerCreateUser)
	app.Post("/v1/api/login", handlers.HandlerUserLogin)
	app.Put("/v1/api/users", middleware.MiddlewareAuth(handlers.HandlerUpdateUser))
	app.Delete("/v1/api/users", middleware.MiddlewareAuth(handlers.HandlerDeleteUser))
	app.Delete("/v1/api/feeds", middleware.MiddlewareAuth(handlers.HandlerDeleteFeed))
	app.Delete("/v1/api/feeds", middleware.MiddlewareAuth(handlers.HandlerDeleteFeedFollow))
}
