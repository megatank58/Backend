package main

import (
	"github.com/gominima/cors"
	"github.com/gominima/middlewares"
	"github.com/gominima/minima"
	_ "github.com/joho/godotenv/autoload"
	"github.com/megatank58/backend/routes"
	"github.com/megatank58/backend/utils/database"
)

func main() {
	database.Setup()

	app := minima.New()
	app.UseRouter(routes.Router())
	app.UseRaw(middleware.Logger)
	app.Use(cors.New().AllowAll())
	app.Listen(":8080")
}
