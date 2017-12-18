package webserver

import (
	"fmt"
	mux "github.com/gorilla/mux"
	cors "github.com/rs/cors"
	"net/http"
	"time"
)

func Server(port int) {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/json")

	r.HandleFunc("/new", Options).Methods("OPTIONS")
	r.HandleFunc("/move", Options).Methods("OPTIONS")
	r.HandleFunc("/move-tile", Options).Methods("OPTIONS")
	r.HandleFunc("/suggest", Options).Methods("OPTIONS")

	r.HandleFunc("/new", New).Methods("GET")
	r.HandleFunc("/move", Move).Methods("POST")
	r.HandleFunc("/move-tile", MoveTile).Methods("POST")
	r.HandleFunc("/suggest", Suggest).Methods("GET")

	allowedOrigins := []string{"*"}
	allowedCredentials := true
	allowedHeaders := []string{
		"Origin",
		"X-Requested-With",
		"Content-Type",
		"Accept",
		"Authorization",
	}
	allowedMethods := []string{
		"POST",
		"GET",
		"PUT",
		"DELETE",
		"PATCH",
		"OPTIONS",
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: allowedCredentials,
		AllowedHeaders:   allowedHeaders,
		AllowedMethods:   allowedMethods,
	})

	server := &http.Server{
		Handler:      c.Handler(r),
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	server.ListenAndServe()
}
