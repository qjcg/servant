// The middleware package provides HTTP server middleware.
package middleware

import (
	"log"
	"net/http"
	"strings"

	"git.jgosset.net/srv/git/color.git"
)

// Log provides middleware for logging HTTP requests.
func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)

		// Extract IPv4 addresses.
		remoteIP := strings.Split(r.RemoteAddr, ":")[0]
		// Extract IPv6 addresses.
		if strings.Contains(r.RemoteAddr, "[") {
			remoteIP = strings.Split(r.RemoteAddr, "]:")[0][1:]
		}

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
			coloredMethod = color.Colored(color.RedB, color.Black, r.Method)
		}

		log.Printf("%-s %-s %s\n",
			coloredMethod,
			r.URL,
			remoteIP,
		)
	})
}
