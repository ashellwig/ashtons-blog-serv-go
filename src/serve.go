package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Argument for where to find the root directory of the site
	selDir := flag.String("dist", "../dist", "destination dist folder")

	flag.Parse()

	// Use arugment from -dist to find root directory
	fs := http.FileServer(http.Dir(*selDir))
	// Serve it at '/'
	http.Handle("/", fs)

	log.Println("Using dist: ", *selDir)
	log.Println("Listening on 127.0.0.1:3000", nil)

	// Start the server
	http.ListenAndServe(":3000", nil)
}
