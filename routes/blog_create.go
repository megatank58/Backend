package routes

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gominima/minima"
	"github.com/megatank58/backend/utils/database"
)

func BlogCreate(res *minima.Response, req *minima.Request) {
	data, _ := ioutil.ReadAll(req.Raw().Body)
	obj := make(map[string]string)
	_ = json.Unmarshal([]byte(data), &obj)
	data, _ = json.Marshal(database.CreateBlog(req.Param("blog"), obj["content"]))

	res.SetHeader("Access-Control-Allow-Origin", "*").OK().Send(string(data))
}
