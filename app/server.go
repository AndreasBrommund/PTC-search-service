package app

import (
	"log"
	"net/http"
	"os"
	"lcd/PTC-search-service/app/web"
)

//StartServer starts the entire web server.
//It is also responsible for initiating the database struct.
func StartServer(port string) {
	router := web.NewRouter(routes)
	router.ServeFiles("/client/*filepath", http.Dir(os.Getenv("PTCFRONTEND")))
	log.Println("starting the webserver...", "http://localhost"+port)
	http.ListenAndServe(port, router)
}
