package routes

import (
	"github.com/gominima/minima"
)

func ProjectGet(res *minima.Response, req *minima.Request) {
	res.SetHeader("Access-Control-Allow-Origin", "*").OK().WriteBytes(RequestRawGitHub(req.Params["project"] + "/main/README.md"))
}
