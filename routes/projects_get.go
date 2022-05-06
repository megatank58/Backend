package routes

import (
	"github.com/gominima/minima"
)

func ProjectsGet(res *minima.Response, req *minima.Request) {
    res.SetHeader("Access-Control-Allow-Origin", "*").OK().WriteBytes(RequestGitHub("users/megatank58/repos"))
}
