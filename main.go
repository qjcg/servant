// A simple static HTTP server, tailored to my preferences.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dir := flag.String("d", ".", "directory to serve")
	port := flag.String("p", "8080", "TCP port to listen on")
	ip := flag.String("i", "127.0.0.1", "IP Address to listen on")
	flag.Usage = func() {
		fmt.Println("What can I do for you, sir?")
		flag.PrintDefaults()
	}
	flag.Parse()

	log.Printf("Serving %s on http://%s:%s/\n", *dir, *ip, *port)
	log.Fatal(http.ListenAndServe(*ip+":"+*port, http.FileServer(http.Dir(*dir))))
}
