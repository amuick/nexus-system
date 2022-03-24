package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"net/http"
)

var logger logrus.Logger

func main() {
	r := mux.NewRouter()

	logger = *logrus.New()

	r.HandleFunc("/ping", ping).Methods("GET")

	logger.Print("Starting server")
	if err := http.ListenAndServe(":3000", r); err != nil {
		logger.Fatalf("error %v", err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("server up"))
}
