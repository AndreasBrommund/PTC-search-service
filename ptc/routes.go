package ptc

import (
	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

func initRoutes() (routes Routes) {
	middleware := alice.New(context.ClearHandler, loggingHandler)

	routes.Get("api version", "/api/1",
		middleware.ThenFunc(apiVersion))

	return
}
