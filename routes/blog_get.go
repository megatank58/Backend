package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/megatank58/backend/utils/database"
)

func BlogGet(ctx *fiber.Ctx) error {
	data, _ := json.Marshal(database.GetBlog(ctx.Params("blog")))

	return ctx.Status(200).Send(data)
}
