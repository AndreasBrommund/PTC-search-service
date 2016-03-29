package ptc

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
)

//loggingHandler is a middleware that logs the time it takes to
//serve a specific endpoint handler.
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("%s: [%s] %q %v\n",
			context.Get(r, "name"),
			r.Method,
			r.URL.String(),
			t2.Sub(t1))
	})
}
