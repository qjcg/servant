// A simple static HTTP server.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.String("p", "8080", "TCP port")
	ip := flag.String("i", "127.0.0.1", "IP address")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [-p <port>] [-i <ipaddr>] [dir]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	var dir string
	switch d := flag.Arg(0); {
	case d == "":
		dir = "."
	default:
		if _, err := os.Stat(d); os.IsNotExist(err) {
			log.Fatalf("Directory does not exist: %s\n", d)
		}
		dir = d
	}

	log.Printf("Serving %s on http://%s:%s/\n", dir, *ip, *port)
	log.Fatal(http.ListenAndServe(*ip+":"+*port, LogMiddleware(http.FileServer(http.Dir(dir)))))
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)
		log.Printf("%v %v %v\n", r.Method, r.URL, strings.Split(r.RemoteAddr, ":")[0])
	})
}
