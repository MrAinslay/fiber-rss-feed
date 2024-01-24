package handlers

import (
	"fmt"
	"time"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/MrAinslay/fiber-rss-feed/packages/utils"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func HandlerCreateFeed(ctx *fiber.Ctx, usr config.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	type payload struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Name      string    `json:"name"`
		URL       string    `json:"url"`
	}

	params := parameters{}
	if err := ctx.BodyParser(&params); err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	feed, err := config.DBQueris.CreateFeed(ctx.Context(), config.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usr.ID,
		Name:      params.Name,
		Url:       params.URL,
	})
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 200, payload{
		Id:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		URL:       feed.Url,
	})
}