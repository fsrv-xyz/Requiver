package main

import (
	"encoding/json"
	"golang.fsrv.services/jsonstatus"
	"log"
	"net"
	"net/http"
	"fmt"
	"time"
)

// PingHandler Add IP Addresses and the Time to an array (pinged)
func PingHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Print(err)
	}
	test := Address{IP: ip, Time: currentTime.String()}
	for i := 0; i < len(pinged); i++ {
		if pinged[i].IP == test.IP {
			jsonstatus.Status{Message: "Address already stored!", StatusCode: http.StatusNotAcceptable}.Encode(w)
			return
		}
	}
	jsonstatus.Status{Message: "IP Address added.", StatusCode: http.StatusOK}.Encode(w)
	addAddress(ip, currentTime.Format(time.RFC3339))
}

// Ack Remove specific IPAddress from the Array (pinged)
func Ack(w http.ResponseWriter, r *http.Request) {
	addr := r.URL.Path[5:]
	for i := 0; i < len(pinged); i++ {
		if pinged[i].IP == addr {
			a := append(pinged[:i], pinged[i+1:]...)
			pinged = nil
			pinged = a
			jsonstatus.Status{Message: fmt.Sprintf("%v deleted", addr), StatusCode: http.StatusOK}.Encode(w)
			return
		}
	}
	jsonstatus.Status{Message: fmt.Sprintf("%v not Found", addr), StatusCode: http.StatusNotAcceptable}.Encode(w)
}

// Flush Remove all IPAddresses from the Array (pinged)
func Flush(w http.ResponseWriter, _ *http.Request) {
	if pinged != nil {
		pinged = nil
		jsonstatus.Status{Message: "Everything Deleted", StatusCode: http.StatusOK}.Encode(w)
	} else {
		jsonstatus.Status{Message: "No Addresses stored", StatusCode: http.StatusNotAcceptable}.Encode(w)
	}

}

// Status Output: Array of IP Addresses and the Time they pinged the Server
func Status(w http.ResponseWriter, r *http.Request) {
	if pinged != nil {
		err := json.NewEncoder(w).Encode(pinged)
		if err != nil {
			log.Printf("Error occured: %v", err)
			return
		}
	} else {
		jsonstatus.Status{Message: "No Address stored", StatusCode: http.StatusNotAcceptable}.Encode(w)
	}
}
