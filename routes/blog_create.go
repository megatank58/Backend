package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/database"
	"github.com/megatank58/backend/utils/request"
)

func BlogCreate(res *minima.Response, req *minima.Request) {

	isAuthenticated := request.CheckAuthentication(req.GetHeader("token"))

	if !isAuthenticated {
		res.SetHeader("Access-Control-Allow-Origin", "*").Unauthorized().Send("Access token is invalid or not of the user")
		return
	}

	data, _ := ioutil.ReadAll(req.Raw().Body)
	obj := make(map[string]string)
	_ = json.Unmarshal([]byte(data), &obj)
	data, _ = json.Marshal(database.CreateBlog(req.Param("blog"), obj["content"]))

	res.SetHeader("Access-Control-Allow-Origin", "*").OK().Send(string(data))
}
