package handlers

import (
	"fmt"
	"time"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/MrAinslay/fiber-rss-feed/packages/utils"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HandlerCreateUser(ctx *fiber.Ctx) {
	type parameters struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	type payload struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		ApiKey    string    `json:"api_key"`
	}
	params := parameters{}

	if err := ctx.BodyParser(&params); err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
	}

	encrPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), 10)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	usr, err := config.DBQueris.CreateUser(ctx.Context(), config.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Password:  string(encrPass),
	})
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
	}

	utils.RespondWithJSON(ctx, 200, payload{
		Id:        usr.ID,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		Name:      usr.Name,
		ApiKey:    usr.ApiKey,
	})
}

func HandlerGetUserByKey(ctx *fiber.Ctx) {
	type payload struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		ApiKey    string    `json:"api_key"`
	}
	apiKey := ctx.Params("api_key")

	usr, err := config.DBQueris.GetUserById(ctx.Context(), apiKey)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 200, payload{
		Id:        usr.ID,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		Name:      usr.Name,
		ApiKey:    usr.ApiKey,
	})
}

func HandlerUserLogin(ctx *fiber.Ctx) {

}
