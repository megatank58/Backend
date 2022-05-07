package routes

import (
	"encoding/json"

	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/database"
)

func BlogCreate(res *minima.Response, req *minima.Request) {
	data, _ := json.Marshal(database.CreateBlog(req.Param("blog"), req.Param("content")))

	res.SetHeader("Access-Control-Allow-Origin", "*").OK().Send(string(data))
}
