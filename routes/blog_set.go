package routes

import (
	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/database"
)

func BlogSet(res *minima.Response, req *minima.Request) {
    database.SetBlog(req.Param("blog"), req.Param("content"));
    
	res.SetHeader("Access-Control-Allow-Origin", "*").OK().Send("OK")
}