package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/megatank58/backend/utils/database"
)

func BlogsGet(ctx *fiber.Ctx) error {
	data, _ := json.Marshal(database.GetBlogs())
	return ctx.Status(200).Send(data)
}
