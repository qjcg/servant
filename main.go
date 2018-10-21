// Simply serve HTTP.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"git.jgosset.net/srv/git/color.git"
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
	log.Fatal(
		http.ListenAndServe(
			*ip+":"+*port,
			LogMiddleware(
				http.FileServer(
					http.Dir(dir)))))
}

// LogMiddleware is middleware for logging HTTP requests.
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)

		remoteIP := strings.Split(r.RemoteAddr, ":")[0]

		var coloredMethod string
		switch r.Method {
		case "GET":
			coloredMethod = color.Colored(color.GreenB, color.Black, r.Method)
		case "HEAD":
			coloredMethod = color.Colored(color.Yellow, color.Black, r.Method)
		case "DELETE":
			coloredMethod = color.Colored(color.RedB, color.Black, r.Method)
		case "POST":
			coloredMethod = color.Colored(color.CyanB, color.Black, r.Method)
		default:
			coloredMethod = color.Colored(color.White, color.Black, r.Method)
		}

		log.Printf("%-s %-s %s\n",
			coloredMethod,
			r.URL,
			remoteIP,
		)
	})
}
