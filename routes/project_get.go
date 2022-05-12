package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/megatank58/backend/utils/request"
)

func ProjectGet(ctx *fiber.Ctx) error {
	return ctx.Status(200).Send(request.RawGitHub(ctx.Params("project") + "/main/README.md"))
}
