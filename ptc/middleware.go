package ptc

import (
	"log"
	"net/http"
	"time"

	"encoding/json"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
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

// recoverHandler recovers from panics and logs the error to stdout
// Response to the caller will contain a message with the error that made
// service crash.
func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(struct {
					Fel string `json:"error"`
				}{"Something when terrible wrong"})
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

	}
}


// Param is a function that gets the parameter value of a specified
// url key.
func Param(r *http.Request, key string) string {
	ps := context.Get(r, "params").(httprouter.Params)
	return ps.ByName(key)
}
