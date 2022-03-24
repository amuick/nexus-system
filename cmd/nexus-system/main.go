package main

import (
	"github.com/gorilla/mux"

	"log"
	"net/http"
)

var logger log.Logger

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ping", ping).Methods("GET")


	logger.Print("Starting server")
	if err := http.ListenAndServe(":3000", r); err != nil {
		logger.Fatalf("error %v", err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
