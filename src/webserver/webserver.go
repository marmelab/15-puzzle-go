package webserver

import (
	"net/http"
	"time"
	"fmt"
	mux "github.com/gorilla/mux"	
)

func Server(port int) {
	r := mux.NewRouter()
	r.HandleFunc("/new", New).Methods("GET")	
	r.HandleFunc("/move", Move).Methods("POST")
	r.HandleFunc("/suggest", Suggest).Methods("POST")

	server := &http.Server{
		Handler: r,
		Addr: fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}	
	server.ListenAndServe()
}
