package webserver

import (
	"fmt"
	mux "github.com/gorilla/mux"
	"net/http"
	"time"
)

func Server(port int) {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")
	r.HandleFunc("/new", New).Methods("GET")
	r.HandleFunc("/move", Move).Methods("POST")
	r.HandleFunc("/suggest", Suggest).Methods("GET")

	server := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.ListenAndServe()
}
