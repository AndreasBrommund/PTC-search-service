package app

import (
	"lcd/PTC-search-service/app/controller"
	"lcd/PTC-search-service/app/web"

	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

//initRoutes constructs the routes exposed in this application.
//It will also connect the routes with the handler functions,
//optionally setting up middleware is also supported.
func routes() (routes web.Routes) {
	middleware := alice.New(context.ClearHandler, web.LoggingHandler, web.RecoverHandler)

	routes.Get("get hastags", "/api/1/tags",
		middleware.ThenFunc(controller.Tags))
	routes.Get("interval", "/api/1/interval",
		middleware.ThenFunc(controller.Interval))

	return
}
