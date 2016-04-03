package ptc

import (
	"log"
	"net/http"
	"os"
)

var database Database
var elasticSearch Elastic

//StartServer starts the entire web server.
//It is also responsible for initiating the database struct.
func StartServer(port string) {
	var err error
	if database, err = NewDatabase("./config/db.json", "dev"); err != nil {
		log.Fatal(err)
	}

	if elasticSearch, err = NewElastic(); err != nil {
		log.Fatal(err)
	}

	router := NewRouter()
	router.ServeFiles("/client/*filepath", http.Dir(os.Getenv("PTCFRONTEND")))
	log.Println("starting the webserver...", "http://localhost"+port)
	http.ListenAndServe(port, router)
}
