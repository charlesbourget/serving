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
	format = flag.String("f", "html", "Format to use. Either html, json or xml")
	flag.Parse()

	switch *format {
	case "html":
		http.Handle("/", http.FileServer(http.Dir(*dir)))
	case "json", "xml":
		http.HandleFunc("/api/", fileServer)
		http.Handle("/", http.FileServer(http.Dir("./web/prod")))
	default:
		log.Fatal("Format not supported...")
	}

	log.Printf("Serving %s on HTTP port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}
