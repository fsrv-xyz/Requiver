package main

import (
	"net/http"
)

type Address struct {
	IP   string
	Time string
}

var pinged []Address

// addAddress Function to add IPAddress and Time to array
func addAddress(ip, currentTime string) {
	pinged = append(pinged, Address{IP: ip, Time: currentTime})
}

func main() {

	http.HandleFunc("/ping", LoggingMiddleware(PingHandler))
	http.HandleFunc("/status", LoggingMiddleware(Status))
	http.HandleFunc("/ack/", LoggingMiddleware(Ack))
	http.HandleFunc("/flush", LoggingMiddleware(Flush))

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}
