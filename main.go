package main

import (
	"github.com/gominima/cors"
	"github.com/gominima/middlewares"
	"github.com/gominima/minima"
	"github.com/megatank58/backend/routes"
)

func main() {
	app := minima.New()
	app.UseRouter(routes.Router())
	app.UseRaw(middleware.Logger)
	app.Use(cors.New().AllowAll())
	app.Listen(":8080")
}
