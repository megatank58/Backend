package routes

import (
	"encoding/json"

	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/database"
)

func BlogGet(res *minima.Response, req *minima.Request) {
	data, _ := json.Marshal(database.GetBlog(req.Param("blog")))

	res.SetHeader("Access-Control-Allow-Origin", "*").OK().Send(string(data))
}
