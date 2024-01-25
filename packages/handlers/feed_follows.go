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

func HandlerDeleteFeedFollow(ctx *fiber.Ctx, usr config.User) {
	id := ctx.Params("id")
	followUUID, err := uuid.Parse(id)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	follow, err := config.DBQueris.GetFeedById(ctx.Context(), followUUID)
	if err != nil {
		utils.RespondWithErr(ctx, 400, fmt.Sprint(err))
		return
	}

	if follow.UserID != usr.ID {
		utils.RespondWithErr(ctx, 501, "Not authorized")
		return
	}

	config.DBQueris.DeleteFeed(ctx.Context(), followUUID)
	utils.RespondWithJSON(ctx, 201, follow)
}
