package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

//loggingHandler is a middleware that logs the time it takes to
//serve a specific endpoint handler.
func LoggingHandler(next http.Handler) http.Handler {
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
func RecoverHandler(next http.Handler) http.Handler {
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

//Param finds and returns the value from a url parameter
//For example /example?token=123aBc would return 123aBc as a string.
func Param(r *http.Request, key string) (s string, err error) {
	s = r.URL.Query().Get(key)
	if len(s) == 0 {
		return "", errors.New("No parameter found for key: " + key)
	}
	return
}

// NamedParam gets the named parameter value of a specified pattern.
// For example a route like /user/:id, one could access the value of :id.
func NamedParam(r *http.Request, key string) string {
	ps := context.Get(r, "params").(httprouter.Params)
	return ps.ByName(key)
}
