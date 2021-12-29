package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var dir *string
var format *string

func main() {
	port := flag.String("p", "8100", "Port to serve on")
	dir = flag.String("d", ".", "The directory of static file to host")
	format = flag.String("f", "plain", "Format to use. Either plain or json")
	flag.Parse()

	switch *format {
	case "plain":
		http.Handle("/", http.FileServer(http.Dir(*dir)))
	case "json", "xml":
		http.HandleFunc("/", fileServer)
	default:
		log.Fatal("Format not supported...")
	}

	log.Printf("Serving %s on HTTP port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}
