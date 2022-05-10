package routes

import (
	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/database"
	"github.com/megatank58/backend/utils/request"
)

func BlogDelete(res *minima.Response, req *minima.Request) {

	isAuthenticated := request.CheckAuthentication(req.GetHeader("token"))

	if !isAuthenticated {
		res.SetHeader("Access-Control-Allow-Origin", "*").Unauthorized().Send("Access token is invalid or not of the user")
		return
	}
	
	database.DeleteBlog(req.Param("blog"))

	res.SetHeader("Access-Control-Allow-Origin", "*").OK().Send("OK")
}
