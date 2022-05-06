package routes

import (
	"github.com/gominima/minima"
)

func Redirect(res *minima.Response, req *minima.Request) {
	res.Redirect("/projects")
}
