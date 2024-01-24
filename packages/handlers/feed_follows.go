package handlers

import (
	"fmt"
	"time"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/MrAinslay/fiber-rss-feed/packages/utils"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func HandlerCreateFeedFollow(ctx *fiber.Ctx, usr config.User) {
	id := ctx.Params("id")
	feedUuid, err := uuid.Parse(id)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	feedFollows, err := config.DBQueris.CreateFeedFollow(ctx.Context(), config.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usr.ID,
		FeedID:    feedUuid,
	})
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 200, feedFollows)
}
