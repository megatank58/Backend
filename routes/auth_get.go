package routes

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/megatank58/backend/utils/request"
)

func AuthGet(ctx *fiber.Ctx) error {
	
	data := request.Oauth("access_token?client_id="+os.Getenv("CLIENT_ID")+"&client_secret="+os.Getenv("CLIENT_SECRET")+"&code="+ctx.Params("code"))

	return ctx.Status(200).Send(data)
}
