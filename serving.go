package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
)

//go:embed static
var static embed.FS

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
		http.Handle("/", http.FileServer(http.FS(static)))
	default:
		log.Fatal("Format not supported...")
	}

	log.Printf("Serving %s on HTTP port: %s\n", *dir, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}
