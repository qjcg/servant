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
	// defaults
	dir, port := ".", "8080"

	switch len(os.Args) {
	case 2:
		if os.Args[1] == "-h" {
			fmt.Println(Usage)
			os.Exit(1)
		} else {
			dir = os.Args[1]
		}
	case 3:
		dir, port = os.Args[1], os.Args[2]
	}

	log.Printf("Serving %s on http://0.0.0.0:%s/\n", dir, port)
	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir(dir))))
}
