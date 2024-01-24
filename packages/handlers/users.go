package handlers

import (
	"encoding/json"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/gofiber/fiber"
)

func HandlerCreateUser(ctx *fiber.Ctx) {
	decoder := json.NewDecoder()

	config.DBQueris.CreateUser(ctx.Context(), config.CreateUserParams{})
}

func HandlerGetUserById(ctx *fiber.Ctx) {

}

func HandlerUserLogin(ctx *fiber.Ctx) {

}
