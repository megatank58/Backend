package routes

import "github.com/gominima/minima"

func Router() *minima.Router {
	router := minima.NewRouter()
	router.Get("/", Redirect)
	router.Get("/projects", ProjectsGet)
	router.Get("/blogs", BlogsGet)
	router.Get("/projects/:project", ProjectGet)
	router.Get("/blogs/:blog", BlogGet)
	router.Get("/blogs/:blog/create/:content", BlogCreate)
	router.Get("/blogs/:blog/set/:content", BlogSet)
	router.Get("/blogs/:blog/delete", BlogDelete)
	return router
}
