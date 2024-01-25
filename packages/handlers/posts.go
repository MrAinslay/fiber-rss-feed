package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/MrAinslay/fiber-rss-feed/packages/config"
	"github.com/MrAinslay/fiber-rss-feed/packages/utils"
	"github.com/gofiber/fiber"
)

func HandlerGetPostsByUser(ctx *fiber.Ctx, usr config.User) {
	limitQry := ctx.Query("limit")
	limit := 10

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
