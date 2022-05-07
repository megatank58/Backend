package routes

import (
	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/request"
)

func ProjectGet(res *minima.Response, req *minima.Request) {
	res.SetHeader("Access-Control-Allow-Origin", "*").OK().WriteBytes(request.RawGitHub(req.Param("project") + "/main/README.md"))
}
