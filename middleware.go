package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

// LoggingMiddleware Middleware function for logging http requests
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	logFunction := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		l := log.New(os.Stdout, timestamp, 0)
		addr, _, _ := net.SplitHostPort(r.RemoteAddr)
		l.Printf(" | %v accessed %v", addr, r.URL.Path)
	}
	return http.HandlerFunc(logFunction)
}
