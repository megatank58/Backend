package routes

import (
	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/database"
)

func BlogDelete(res *minima.Response, req *minima.Request) {
	database.DeleteBlog(req.Param("blog"))

	res.SetHeader("Access-Control-Allow-Origin", "*").OK().Send("OK")
}
