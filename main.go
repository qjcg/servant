package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	dir := flag.String("d", ".", "directory to serve")
	port := flag.String("p", "8080", "TCP port to listen on")
	flag.Parse()

	log.Printf("Serving %s on http://127.0.0.1:%s/\n", *dir, *port)
	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*dir))))
}
