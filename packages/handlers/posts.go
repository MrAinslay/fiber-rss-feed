package handlers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/MrAinslay/fiber-rss-feed/packages/utils"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func HandlerGetPostsByUser(ctx *fiber.Ctx, usr config.User) {
	type payload struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UserId    uuid.UUID `json:"user_id"`
		FeedId    uuid.UUID `json:"feed_id`
	}

	limitQry := ctx.Query("limit")
	limit := 10

	log.Println(int32(limit))

	var err error
	if limitQry != "" {
		limit, err = strconv.Atoi(limitQry)
		if err != nil {
			log.Printf("Error formating limit query defaulting to 10: %v", err)
			limit = 10
		}
	}

	posts, err := config.DBQueris.GetPostsByUser(ctx.Context(), config.GetPostsByUserParams{
		UserID: usr.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		utils.RespondWithErr(ctx, 401, fmt.Sprint(err))
		return
	}

	utils.RespondWithJSON(ctx, 201, posts)
}
