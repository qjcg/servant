// The middleware package provides HTTP server middleware.
package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

var (
	green  = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).Add(color.Bold).SprintFunc()
	cyan   = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
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
			coloredMethod = green(r.Method)
		case "HEAD":
			coloredMethod = yellow(r.Method)
		case "DELETE":
			coloredMethod = red(r.Method)
		case "POST":
			coloredMethod = cyan(r.Method)
		default:
			coloredMethod = red(r.Method)
		}

		log.Printf("%-s %-s %s\n",
			coloredMethod,
			r.URL,
			remoteIP,
		)
	})
}
