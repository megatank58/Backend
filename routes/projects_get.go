package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/megatank58/backend/utils/request"
)

func ProjectsGet(ctx *fiber.Ctx) error {
	return ctx.Status(200).Send(request.GitHub("users/megatank58/repos"))
}
