package routes

import (
	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/request"
)

func ProjectsGet(res *minima.Response, req *minima.Request) {
    res.SetHeader("Access-Control-Allow-Origin", "*").OK().WriteBytes(request.GitHub("users/megatank58/repos"))
}
