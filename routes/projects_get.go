package routes

import (
	"github.com/gominima/minima"
)

func ProjectsGet(res *minima.Response, req *minima.Request) {
    res.OK().WriteBytes(Request("users/megatank58/repos"))
}
