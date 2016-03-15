package main

import (
	"flag"
	"fmt"

	"github.com/lcd/PTC-search-service/ptc"
)

//Entry point of Search service application
func main() {
	port := flag.String("port", ":8080", "The port the webserver should run on.")
	flag.Parse()
	fmt.Println("Port:", *port)
	ptc.StartServer(*port)
}
