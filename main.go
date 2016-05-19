package main

import (
	"flag"
	"lcd/PTC-search-service/app"
	"lcd/PTC-search-service/app/storage"
	"log"
)

//Entry point of Search service application
//Parse command line flag and tries the connection to elastic.
func main() {
	port := flag.String("port", ":8080", "The port the webserver should run on.")
	flag.Parse()
	if err := storage.Connect(); err != nil {
		log.Println("Cannot connect to elastic..")
	} //Init database
	app.StartServer(*port)
}
