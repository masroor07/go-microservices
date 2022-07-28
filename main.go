package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	details "github.com/masroor07/go-microservices/details"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Status is up")
	response := map[string]string{
		"status": "UP",
		"time":   time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Welcome to homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Serving Homepage")
}
func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching details")
	hostName, err := details.GetHostname()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()
	response := map[string]string{
		"hostname": hostName,
		"ip":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Println("Server has started")

	http.ListenAndServe(":80", r)
}
