package ptc

import (
	"log"
	"net/http"
	"reflect"
	"time"

	"encoding/json"
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

// Decodes a request body into the struct passed to the middleware.
// If the request body is not JSON, it will return a 400 Bad Request error.
// Stores the decoded body into a context object.
func jsonParserHandler(v interface{}) func(http.Handler) http.Handler {
	t := reflect.TypeOf(v)

	m := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(t).Interface()
			err := json.NewDecoder(r.Body).Decode(val)

			if err != nil {
				log.Println(err)
				panic(err)
				return
			}

			if next != nil {
				context.Set(r, "body", val)
				next.ServeHTTP(w, r)
			}
		}
		return http.HandlerFunc(fn)
	}
	return m
}
