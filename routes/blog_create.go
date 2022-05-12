package routes

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/megatank58/backend/utils/database"
	"github.com/megatank58/backend/utils/request"
)

func BlogCreate(ctx *fiber.Ctx) error {

	isAuthenticated := request.CheckAuthentication(ctx.GetReqHeaders()["token"])

	if !isAuthenticated {
		return ctx.Status(401).SendString("Access token is invalid or not of the user")
	}
	
	payload := database.Blog{}

    if err := ctx.BodyParser(&payload); err != nil {
        return err
    }

	data, _ := json.Marshal(database.CreateBlog(ctx.Params("blog"), payload.Content))

	return ctx.Status(200).Send(data)
}
