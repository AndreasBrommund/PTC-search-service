package app

import (
	"lcd/PTC-search-service/app/web"
	"log"
	"net/http"
	"os"
)

//StartServer starts the entire web server.
//Connects the router with the routes and
//starts the file hosting of the frontend
func StartServer(port string) {
	router := web.NewRouter(routes)
	router.ServeFiles("/client/*filepath", http.Dir(os.Getenv("PTCFRONTEND")))
	log.Println("starting the webserver...", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, router))
}
