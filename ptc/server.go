package ptc

import (
	"log"

	"net/http"

	"github.com/lcd/PTC-search-service/db"
)

var database db.Database

func StartServer(port string) {
	var err error
	database, err = db.NewDatabase("./config/db.json", "dev")
	if err != nil {
		log.Fatal(err)
	}
	router := NewRouter()
	log.Println("starting the webserver...", "http://localhost"+port)
	http.ListenAndServe(port, router)
}
