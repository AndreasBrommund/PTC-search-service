package ptc

import (
	"log"
	"net/http"
)

func StartServer(port string) {
	router := NewRouter()
	log.Println("starting the webserver...", "http://localhost"+port)
	http.ListenAndServe(port, router)
}
