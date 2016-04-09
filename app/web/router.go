package web

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

//Routes is a wrapper around the Route struct
type Routes []Route

//Router is a struct that represents routes in the
//application. It contains a name, method, path and a function that
//handles the route.
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.Handler
}

//Get convenient function that creates Get route and appends it to
//the Routes instance
func (this *Routes) Get(name, path string, handlerFunc http.Handler) {
	*this = append(*this, Route{name, "GET", path, handlerFunc})
}

//Post convenient function that creates a POST router and appends it to
//the Routes instance.
func (this *Routes) Post(name, path string, handlerFunc http.Handler) {
	*this = append(*this, Route{name, "POST", path, handlerFunc})
}

//NewRouter creates a new httprouter.Router that have all the
//routes defined under the initRoutes function.
func NewRouter(initRoutes func() Routes) *httprouter.Router {
	router := httprouter.New()
	for _, route := range initRoutes() {
		router.Handle(route.Method, route.Path, wrapHandler(route.HandlerFunc, route.Name))
	}
	return router
}

//wrapHandler is a wrapper around the ServerHTTP interface to
//make httprouter.routes adhere to the interface. Conveniently adds the
//route parameters and the route name to the context.
func wrapHandler(h http.Handler, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		context.Set(r, "name", name)
		h.ServeHTTP(w, r)
	}
}
