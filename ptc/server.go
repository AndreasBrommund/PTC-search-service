package ptc

import (
	"log"

	"net/http"

	"os"

	"github.com/lcd/PTC-search-service/db"
)

var database db.Database

//StartServer starts the entire web server.
//It is also responsible for initiating the database struct.
func StartServer(port string) {
	var err error
	database, err = db.NewDatabase("./config/db.json", "dev")
	if err != nil {
		log.Fatal(err)
	}
	router := NewRouter()
	router.ServeFiles("/client/*filepath", http.Dir(os.Getenv("PTCFRONTEND")))
	log.Println("starting the webserver...", "http://localhost"+port)
	http.ListenAndServe(port, router)
}
