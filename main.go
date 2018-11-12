// Simply serve HTTP.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/qjcg/servant/pkg/middleware"
)

func main() {
	tlsCert := flag.String("c", "", "TLS certificate")
	tlsKey := flag.String("k", "", "TLS key")
	port := flag.String("p", "8080", "TCP port")
	ip := flag.String("i", "127.0.0.1", "IP address")

	flag.Usage = func() {
		fmt.Printf("Usage: %s [-c <TLSCert>] [-k <TLSKey>] [-i <IPAddr>] [-p <Port>] [Dir]\n", os.Args[0])
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

	// If a TLS key & cert are provided, serve HTTPS, otherwise HTTP.
	if *tlsCert != "" && *tlsKey != "" {
		log.Printf("Serving %s on https://%s:%s/\n", dir, *ip, *port)
		log.Fatal(
			http.ListenAndServeTLS(
				*ip+":"+*port,
				*tlsCert, *tlsKey,
				middleware.Log(
					http.FileServer(
						http.Dir(dir)))))
	} else {
		log.Printf("Serving %s on http://%s:%s/\n", dir, *ip, *port)
		log.Fatal(
			http.ListenAndServe(
				*ip+":"+*port,
				middleware.Log(
					http.FileServer(
						http.Dir(dir)))))
	}
}
