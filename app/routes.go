package app

import (
	"github.com/gorilla/context"
	"github.com/justinas/alice"
	"lcd/PTC-search-service/app/controller"
"lcd/PTC-search-service/app/web"
)

//initRoutes constructs the routes exposed in this application.
//It will also connect the routes with the handler functions,
//optionally setting up middleware is also supported.
func routes() (routes web.Routes) {
	middleware := alice.New(context.ClearHandler, web.LoggingHandler,web.RecoverHandler)

	routes.Get("api version", "/api/1/",
		middleware.ThenFunc(controller.ApiVersion))


	routes.Get("get hastags", "/api/1/tags",
		middleware.ThenFunc(controller.GetHastags))


	routes.Get("getting hashtags for user", "/api/1/hashtags",
		middleware.ThenFunc(controller.GetTweetsFromUserID))

	return
}
