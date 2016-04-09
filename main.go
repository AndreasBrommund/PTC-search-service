package main

import (
	"flag"
	"fmt"
	"lcd/PTC-search-service/app"
	"lcd/PTC-search-service/app/storage"
	"log"
)

//Entry point of Search service application
func main() {
	port := flag.String("port", ":8080", "The port the webserver should run on.")
	flag.Parse()
	fmt.Println("Port:", *port)
	if err := storage.Connect(); err != nil {
		log.Println("Cannot connect to elastic..")
	} //Init database
	app.StartServer(*port)
}
