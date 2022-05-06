package routes

import (
	"io/ioutil"
	"net/http"
	"github.com/gominima/minima"
)

func Router() *minima.Router {
	router := minima.NewRouter()
	router.Get("/", Redirect)
	router.Get("/projects", ProjectsGet)
	router.Get("/blogs", BlogsGet)
	router.Get("/projects/:project", ProjectGet)
	router.Get("/blogs/:blog", BlogGet)
	return router
}

func Request(route string) []byte {
	response, err := http.Get("https://api.github.com/" + route)
    if err != nil {
        panic(err)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        panic(err)
    }
	return responseData
}
