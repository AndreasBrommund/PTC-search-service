package ptc

import (
	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

//initRoutes constructs the routes exposed in this application.
//It will also connect the routes with the handler functions,
//optionally setting up middleware is also supported.
func initRoutes() (routes Routes) {
	middleware := alice.New(context.ClearHandler, loggingHandler)

	routes.Get("api version", "/api/1",
		middleware.ThenFunc(apiVersion))

	routes.Get("Tweet count", "/api/1/count/tweet",
		middleware.ThenFunc(countTweet))
	return
}
