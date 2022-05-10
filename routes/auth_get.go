package routes

import (
	"os"

	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/request"
)

func AuthGet(res *minima.Response, req *minima.Request) {
	
	data := request.Oauth("access_token?client_id="+os.Getenv("CLIENT_ID")+"&client_secret="+os.Getenv("CLIENT_SECRET")+"&code="+req.Param("code"))

	res.SetHeader("Access-Control-Allow-Origin", "*").OK().WriteBytes(data)
}
