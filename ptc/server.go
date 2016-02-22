package ptc

import (
	"log"
	"net/http"
)

func StartServer(port string) {
	log.Println("starting the webserver...", "http://localhost"+port)
	http.ListenAndServe(port, nil)
}
