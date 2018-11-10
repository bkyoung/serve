package main

import (
	"flag"
	"log"
	"net/http"
)

// Because sometimes you just need to quickly share files over http
func main() {
	port := flag.String("p", "8080", "http port to serve on")
	docroot := flag.String("d", ".", "the directory to serve")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*docroot)))

	log.Printf("Serving %s on HTTP port: %s\n", *docroot, *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}
