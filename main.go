// A simple static HTTP server, tailored to my preferences.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const Usage string = `Usage: servant [DIR] [PORT]`

func main() {
	dir := "."
	port := "8080"

	switch {
	case len(os.Args) == 2 && os.Args[1] == "-h":
		fmt.Println(Usage)
		os.Exit(0)
	case len(os.Args) == 2:
		dir = os.Args[1]
	case len(os.Args) == 3:
		dir, port = os.Args[1], os.Args[2]
	}

	log.Printf("Serving %s on http://127.0.0.1:%s/\n", dir, port)
	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir(dir))))
}
